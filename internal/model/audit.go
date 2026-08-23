package model

import "time"

type AuditKind string

const (
	AuditSampleReceived  AuditKind = "sample_received"
	AuditSampleMoved     AuditKind = "sample_moved"
	AuditAssayQueued     AuditKind = "assay_queued"
	AuditAssayStarted    AuditKind = "assay_started"
	AuditResultReceived  AuditKind = "result_received"
	AuditReviewCompleted AuditKind = "review_completed"
	AuditExceptionRaised AuditKind = "exception_raised"
)

type AuditEntry struct {
	ID        string
	Kind      AuditKind
	SubjectID string
	Actor     string
	Location  string
	At        time.Time
	Metadata  map[string]string
}

func (a AuditEntry) IsValid() bool {
	return a.ID != "" && a.Kind != "" && a.SubjectID != "" && a.Actor != "" && !a.At.IsZero()
}

type Exception struct {
	ID        string
	SubjectID string
	Code      string
	Message   string
	Severity  string
	OpenedAt  time.Time
	ClosedAt  *time.Time
}

func (e Exception) Open() bool          { return e.ClosedAt == nil }
func (e *Exception) Close(at time.Time) { e.ClosedAt = &at }
