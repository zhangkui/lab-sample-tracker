package domain

import "time"

type Assay17 struct {
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

func (a Assay17) Start(at time.Time) Assay17 {
	a.StartedAt = at
	return a
}

func (a Assay17) Complete(result string, at time.Time) Assay17 {
	a.Result = result
	a.FinishedAt = at
	return a
}

func (a Assay17) Review(quality string) Assay17 {
	a.Quality = quality
	a.Reviewed = true
	return a
}

func (a Assay17) IsComplete() bool {
	return a.ID != "" && a.SampleID != "" && !a.StartedAt.IsZero() && !a.FinishedAt.IsZero() && a.Result != ""
}

func (a Assay17) NeedsReview() bool {
	return a.IsComplete() && !a.Reviewed
}
