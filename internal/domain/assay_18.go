package domain

import "time"

type Assay18 struct {
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

func (a Assay18) Start(at time.Time) Assay18 {
	a.StartedAt = at
	return a
}

func (a Assay18) Complete(result string, at time.Time) Assay18 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay18) Review(quality string) Assay18 {
	a.Quality = quality
	a.Reviewed = true
	return a
}

func (a Assay18) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}

func (a Assay18) NeedsReview() bool {
	return a.IsComplete() && !a.Reviewed
}
