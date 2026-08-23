package model

import "time"

type WorkflowSnapshot struct {
	Sample     Sample
	Batch      BatchRecord
	Assays     []AssayRequest
	Results    []InstrumentResult
	Audit      []AuditEntry
	Exceptions []Exception
	CapturedAt time.Time
}

func (w WorkflowSnapshot) HasOpenExceptions() bool {
	for _, e := range w.Exceptions {
		if e.Open() {
			return true
		}
	}
	return false
}
func (w WorkflowSnapshot) ApprovedAssays() int {
	n := 0
	for _, a := range w.Assays {
		if a.State == AssayApproved {
			n++
		}
	}
	return n
}
func (w WorkflowSnapshot) ReadyToRelease() bool {
	return w.Batch.State == BatchInReview && !w.HasOpenExceptions() && w.ApprovedAssays() == len(w.Assays)
}
