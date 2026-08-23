package domain

import "time"

type Assay16 struct {
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

func (a Assay16) Start(at time.Time) Assay16 {
	a.StartedAt = at
	return a
}

func (a Assay16) Complete(result string, at time.Time) Assay16 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay16) Review(quality string) Assay16 {
	a.Quality = quality
	a.Reviewed = true
	return a
}

func (a Assay16) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}

func (a Assay16) NeedsReview() bool {
	return a.IsComplete() && !a.Reviewed
}
