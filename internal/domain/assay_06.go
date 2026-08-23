package domain

import "time"

type Assay6 struct {
	ID         string
	SampleID   string
	Name       string
	Instrument string
	StartedAt  time.Time
	FinishedAt time.Time
	Result     string
	Notes      string
	Analyst    string
	Reviewed   bool
}

func (a Assay6) Start(at time.Time) Assay6 {
	a.StartedAt = at
	return a
}

func (a Assay6) Complete(result string, at time.Time) Assay6 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay6) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}

func (a Assay6) NeedsReview() bool {
	return a.IsComplete() && !a.Reviewed
}
