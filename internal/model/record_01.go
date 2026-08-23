package model

import "time"

type ChainRecord1 struct {
	ID         string
	SampleID   string
	Action     string
	Actor      string
	Location   string
	Instrument string
	CreatedAt  time.Time
	Verified   bool
	Notes      string
}

func (r ChainRecord1) Complete() bool {
	return r.ID != "" && r.SampleID != "" && r.Action != "" && r.Verified
}

func (r ChainRecord1) Age(now time.Time) time.Duration {
	if r.CreatedAt.IsZero() {
		return 0
	}
	return now.Sub(r.CreatedAt)
}

func (r ChainRecord1) WithNote(note string) ChainRecord1 {
	r.Notes = note
	return r
}
