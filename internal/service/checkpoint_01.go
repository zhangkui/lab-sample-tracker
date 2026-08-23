package service

import (
	"sort"
	"time"
)

type Checkpoint1 struct {
	Name        string
	Required    bool
	Completed   bool
	CompletedAt time.Time
}

func (c Checkpoint1) Ready() bool {
	return !c.Required || c.Completed
}

func Checkpoints1() []Checkpoint1 {
	out := []Checkpoint1{
		{Name: "intake", Required: true},
		{Name: "aliquot", Required: true},
		{Name: "assay", Required: true},
		{Name: "archive", Required: false},
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func CompleteCheckpoint1(c Checkpoint1, at time.Time) Checkpoint1 {
	c.Completed = true
	c.CompletedAt = at
	return c
}
