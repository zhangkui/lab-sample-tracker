package service

import (
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"sort"
)

type Report struct {
	Total    int
	ByStatus map[model.SampleStatus]int
	Latest   []model.Sample
}

func (l *Lab) Report() Report {
	items := l.List("", 0, 1000)
	r := Report{Total: len(items), ByStatus: map[model.SampleStatus]int{}}
	for _, v := range items {
		r.ByStatus[v.Status]++
	}
	sort.Slice(items, func(i, j int) bool { return items[i].UpdatedAt.After(items[j].UpdatedAt) })
	if len(items) > 5 {
		items = items[:5]
	}
	r.Latest = items
	return r
}
