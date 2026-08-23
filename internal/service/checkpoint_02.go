package service

import (
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"sort"
	"time"
)

type Checkpoint2 struct {
	Name        string
	Required    bool
	Completed   bool
	CompletedAt time.Time
}

func (c Checkpoint2) Ready() bool {
	return !c.Required || c.Completed
}

func Checkpoints2() []Checkpoint2 {
	out := []Checkpoint2{
		{Name: "intake", Required: true},
		{Name: "aliquot", Required: true},
		{Name: "assay", Required: true},
		{Name: "archive", Required: false},
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func CompleteCheckpoint2(c Checkpoint2, at time.Time) Checkpoint2 {
	c.Completed = true
	c.CompletedAt = at
	return c
}
