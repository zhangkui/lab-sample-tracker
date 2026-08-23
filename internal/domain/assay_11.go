package domain

import "time"

type Assay11 struct {
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

func (a Assay11) Start(at time.Time) Assay11 {
	a.StartedAt = at
	return a
}

func (a Assay11) Complete(result string, at time.Time) Assay11 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay11) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}

func (a Assay11) NeedsReview() bool {
	return a.IsComplete() && !a.Reviewed
}
