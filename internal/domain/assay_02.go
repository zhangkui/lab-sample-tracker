package domain

import "time"

type Assay2 struct {
	ID         string
	SampleID   string
	Name       string
	Instrument string
	StartedAt  time.Time
	FinishedAt time.Time
	Result     string
	Notes      string
}

func (a Assay2) Complete(result string, at time.Time) Assay2 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay2) IsComplete() bool {
	return !a.FinishedAt.IsZero() && a.Result != ""
}
