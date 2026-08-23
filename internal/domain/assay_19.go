package domain

import "time"

type Assay19 struct {
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

func (a Assay19) Start(at time.Time) Assay19 {
	a.StartedAt = at
	return a
}

func (a Assay19) Complete(result string, at time.Time) Assay19 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay19) Review(quality string) Assay19 {
	a.Quality = quality
	a.Reviewed = true
	return a
}

func (a Assay19) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}

func (a Assay19) NeedsReview() bool {
	return a.IsComplete() && !a.Reviewed
}
