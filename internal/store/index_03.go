package store

import (
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"strings"
)

type Index3 struct {
	Prefix string
	Values map[string]string
}

func NewIndex3(prefix string) Index3 {
	return Index3{Prefix: strings.TrimSpace(prefix), Values: map[string]string{}}
}

func (i Index3) Key(value string) string {
	return i.Prefix + ":" + strings.TrimSpace(value)
}

func (i Index3) Match(sample model.Sample) bool {
	return strings.HasPrefix(sample.ID, i.Prefix)
}
