package store

import (
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"strings"
)

type Index1 struct {
	Prefix string
	Values map[string]string
}

func NewIndex1(prefix string) Index1 {
	return Index1{Prefix: strings.TrimSpace(prefix), Values: map[string]string{}}
}

func (i Index1) Key(value string) string {
	return i.Prefix + ":" + strings.TrimSpace(value)
}

func (i Index1) Match(sample model.Sample) bool {
	return strings.HasPrefix(sample.ID, i.Prefix)
}
