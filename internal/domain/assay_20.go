package domain

import "time"

type Assay20 struct {
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
	Quality    string
}

func (a Assay20) Start(at time.Time) Assay20 {
	a.StartedAt = at
	return a
}

func (a Assay20) Complete(result string, at time.Time) Assay20 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay20) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}
