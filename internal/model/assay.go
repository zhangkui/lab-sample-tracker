package model

import (
	"errors"
	"strings"
	"time"
)

type AssayState string

const (
	AssayQueued         AssayState = "queued"
	AssayRunning        AssayState = "running"
	AssayAwaitingReview AssayState = "awaiting_review"
	AssayApproved       AssayState = "approved"
	AssayRejected       AssayState = "rejected"
)

type AssayRequest struct {
	ID              string
	SampleID        string
	BatchID         string
	ProtocolID      string
	ProtocolVersion int
	InstrumentID    string
	State           AssayState
	RequestedBy     string
	RequestedAt     time.Time
	StartedAt       *time.Time
	FinishedAt      *time.Time
	ResultID        string
}

func (a AssayRequest) ReadyForExecution() bool {
	return a.State == AssayQueued && a.SampleID != "" && a.ProtocolID != "" && a.InstrumentID != ""
}

func (a *AssayRequest) Start(at time.Time) error {
	if !a.ReadyForExecution() {
		return errors.New("assay is not ready")
	}
	a.State = AssayRunning
	a.StartedAt = &at
	return nil
}

func (a *AssayRequest) Finish(resultID string, at time.Time) error {
	if a.State != AssayRunning || strings.TrimSpace(resultID) == "" {
		return errors.New("assay cannot be finished")
	}
	a.State = AssayAwaitingReview
	a.ResultID = resultID
	a.FinishedAt = &at
	return nil
}

func (a *AssayRequest) Approve() error {
	if a.State != AssayAwaitingReview {
		return errors.New("assay is not awaiting review")
	}
	a.State = AssayApproved
	return nil
}

func (a *AssayRequest) Reject() error {
	if a.State != AssayAwaitingReview {
		return errors.New("assay is not awaiting review")
	}
	a.State = AssayRejected
	return nil
}
