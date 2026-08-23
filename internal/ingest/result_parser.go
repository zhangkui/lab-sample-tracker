package ingest

import (
	"encoding/json"
	"errors"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"strings"
	"time"
)

type WireResult struct {
	RunID        string    `json:"run_id"`
	AssayID      string    `json:"assay_id"`
	InstrumentID string    `json:"instrument_id"`
	ReceivedAt   time.Time `json:"received_at"`
	Readings     []struct {
		Name  string  `json:"name"`
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"readings"`
}

func ParseResult(raw []byte) (model.InstrumentResult, error) {
	var w WireResult
	if len(strings.TrimSpace(string(raw))) == 0 {
		return model.InstrumentResult{}, errors.New("empty instrument payload")
	}
	if err := json.Unmarshal(raw, &w); err != nil {
		return model.InstrumentResult{}, err
	}
	if w.RunID == "" || w.AssayID == "" || w.InstrumentID == "" {
		return model.InstrumentResult{}, errors.New("instrument result identity is incomplete")
	}
	r := model.InstrumentResult{ID: w.RunID + ":" + w.AssayID, RunID: w.RunID, AssayID: w.AssayID, InstrumentID: w.InstrumentID, RawPayload: string(raw), ReceivedAt: w.ReceivedAt}
	for _, x := range w.Readings {
		r.Measurements = append(r.Measurements, model.Measurement{Name: x.Name, Value: x.Value, Unit: x.Unit})
	}
	r.Seal()
	return r, nil
}
