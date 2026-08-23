package model

import "time"

type ChainRecord3 struct {
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

func (r ChainRecord3) Complete() bool {
	return r.ID != "" && r.SampleID != "" && r.Action != "" && r.Verified
}

func (r ChainRecord3) Age(now time.Time) time.Duration {
	if r.CreatedAt.IsZero() {
		return 0
	}
	return now.Sub(r.CreatedAt)
}

func (r ChainRecord3) WithNote(note string) ChainRecord3 {
	r.Notes = note
	return r
}
