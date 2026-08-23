package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/zhangkui/lab-sample-tracker/internal/model"
)

type AssayWorkflow struct {
	lab         *Lab
	protocols   *ProtocolCatalog
	assays      map[string]*model.AssayRequest
	instruments map[string]*model.Instrument
	results     map[string]model.InstrumentResult
	mu          sync.RWMutex
}

func (l *Lab) NewAssayWorkflow() *AssayWorkflow {
	return &AssayWorkflow{lab: l, protocols: NewProtocolCatalog(), assays: map[string]*model.AssayRequest{}, instruments: map[string]*model.Instrument{}, results: map[string]model.InstrumentResult{}}
}
func (w *AssayWorkflow) RegisterProtocol(p model.Protocol) error { return w.protocols.Put(p) }
func (w *AssayWorkflow) RegisterInstrument(i model.Instrument) error {
	if i.ID == "" {
		return errors.New("instrument id is required")
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, ok := w.instruments[i.ID]; ok {
		return errors.New("instrument already registered")
	}
	w.instruments[i.ID] = &i
	return nil
}
func (w *AssayWorkflow) Queue(a model.AssayRequest) error {
	if a.ID == "" || a.SampleID == "" || a.ProtocolID == "" || a.InstrumentID == "" {
		return errors.New("assay fields are incomplete")
	}
	protocol, err := w.protocols.Get(a.ProtocolID)
	if err != nil {
		return fmt.Errorf("validate protocol: %w", err)
	}
	a.ProtocolVersion = protocol.Version
	a.State = model.AssayQueued
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, ok := w.assays[a.ID]; ok {
		return errors.New("assay already exists")
	}
	w.assays[a.ID] = &a
	return w.lab.record(a.SampleID, "assay_queued", map[string]string{"assay_id": a.ID, "protocol_version": fmt.Sprint(a.ProtocolVersion)})
}
func (w *AssayWorkflow) Start(ctx context.Context, id string, at time.Time) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	a, ok := w.assays[id]
	if !ok {
		return errors.New("assay not found")
	}
	i, ok := w.instruments[a.InstrumentID]
	if !ok {
		return errors.New("instrument not found")
	}
	if !i.CanRun(a.ProtocolID) {
		return errors.New("instrument cannot run protocol")
	}
	if err := i.Begin(id); err != nil {
		return err
	}
	if err := a.Start(at); err != nil {
		_ = i.End(id)
		return err
	}
	return w.lab.record(a.SampleID, "assay_started", map[string]string{"assay_id": id})
}
func (w *AssayWorkflow) ReceiveResult(ctx context.Context, r model.InstrumentResult) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := r.Validate(); err != nil {
		return err
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, exists := w.results[r.ID]; exists {
		return errors.New("result already received")
	}
	a, ok := w.assays[r.AssayID]
	if !ok {
		return errors.New("assay not found")
	}
	if a.InstrumentID != r.InstrumentID {
		return errors.New("result instrument does not match assay")
	}
	if a.State != model.AssayRunning {
		return errors.New("assay is not running")
	}
	if err := a.Finish(r.ID, r.ReceivedAt); err != nil {
		return err
	}
	i := w.instruments[r.InstrumentID]
	if i == nil {
		return errors.New("instrument not found")
	}
	if err := i.End(a.ID); err != nil {
		return fmt.Errorf("release instrument: %w", err)
	}
	w.results[r.ID] = r
	return w.lab.record(a.SampleID, "result_received", map[string]string{"result_id": r.ID})
}
func (w *AssayWorkflow) Review(id string, approved bool) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	a, ok := w.assays[id]
	if !ok {
		return errors.New("assay not found")
	}
	var err error
	if approved {
		err = a.Approve()
	} else {
		err = a.Reject()
	}
	if err != nil {
		return err
	}
	return w.lab.record(a.SampleID, "review_completed", map[string]string{"assay_id": id, "approved": fmt.Sprint(approved)})
}
func (w *AssayWorkflow) Assay(id string) (model.AssayRequest, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	a, ok := w.assays[id]
	if !ok {
		return model.AssayRequest{}, errors.New("assay not found")
	}
	return *a, nil
}
func (w *AssayWorkflow) Results() []model.InstrumentResult {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := make([]model.InstrumentResult, 0, len(w.results))
	for _, result := range w.results {
		out = append(out, result)
	}
	return out
}
func (w *AssayWorkflow) AssaysForBatch(batchID string) []model.AssayRequest {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := []model.AssayRequest{}
	for _, assay := range w.assays {
		if assay.BatchID == batchID {
			out = append(out, *assay)
		}
	}
	return out
}
