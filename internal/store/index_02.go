package store

import (
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"strings"
)

type Index2 struct {
	Prefix string
	Values map[string]string
}

func NewIndex2(prefix string) Index2 {
	return Index2{Prefix: strings.TrimSpace(prefix), Values: map[string]string{}}
}

func (i Index2) Key(value string) string {
	return i.Prefix + ":" + strings.TrimSpace(value)
}

func (i Index2) Match(sample model.Sample) bool {
	return strings.HasPrefix(sample.ID, i.Prefix)
}
