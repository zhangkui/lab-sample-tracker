package service

import (
	"errors"
	"fmt"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"sync"
	"time"
)

type Reconciler struct {
	mu         sync.Mutex
	exceptions map[string]model.Exception
}

func NewReconciler() *Reconciler { return &Reconciler{exceptions: map[string]model.Exception{}} }
func (r *Reconciler) Compare(expected []model.AssayRequest, received []model.InstrumentResult, now time.Time) []model.Exception {
	r.mu.Lock()
	defer r.mu.Unlock()
	seen := map[string]bool{}
	for _, result := range received {
		seen[result.AssayID] = true
	}
	out := []model.Exception{}
	for _, assay := range expected {
		if assay.State != model.AssayRejected && !seen[assay.ID] {
			e := model.Exception{ID: fmt.Sprintf("missing-%s", assay.ID), SubjectID: assay.SampleID, Code: "missing_result", Message: "approved assay has no result", Severity: "high", OpenedAt: now}
			r.exceptions[e.ID] = e
			out = append(out, e)
		}
	}
	return out
}
func (r *Reconciler) Open(subject string) []model.Exception {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := []model.Exception{}
	for _, e := range r.exceptions {
		if e.SubjectID == subject && e.Open() {
			out = append(out, e)
		}
	}
	return out
}
func (r *Reconciler) Close(id string, at time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	e, ok := r.exceptions[id]
	if !ok {
		return errors.New("exception not found")
	}
	if !e.Open() {
		return errors.New("exception is closed")
	}
	e.Close(at)
	r.exceptions[id] = e
	return nil
}
