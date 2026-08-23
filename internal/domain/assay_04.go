package domain

import "time"

type Assay4 struct {
	ID         string
	SampleID   string
	Name       string
	Instrument string
	StartedAt  time.Time
	FinishedAt time.Time
	Result     string
	Notes      string
}

func (a Assay4) Complete(result string, at time.Time) Assay4 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay4) IsComplete() bool {
	return !a.FinishedAt.IsZero() && a.Result != ""
}
