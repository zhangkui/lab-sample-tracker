package domain

import "time"

type Assay1 struct {
	ID         string
	SampleID   string
	Name       string
	Instrument string
	StartedAt  time.Time
	FinishedAt time.Time
	Result     string
	Notes      string
}

func (a Assay1) Complete(result string, at time.Time) Assay1 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay1) IsComplete() bool {
	return !a.FinishedAt.IsZero() && a.Result != ""
}
