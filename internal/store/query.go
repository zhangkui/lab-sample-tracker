package store

import (
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"sort"
)

func SortByUpdate(items []model.Sample) []model.Sample {
	sort.Slice(items, func(i, j int) bool { return items[i].UpdatedAt.Before(items[j].UpdatedAt) })
	return items
}
func FilterStatus(items []model.Sample, status model.SampleStatus) []model.Sample {
	out := make([]model.Sample, 0)
	for _, v := range items {
		if status == "" || v.Status == status {
			out = append(out, v)
		}
	}
	return out
}
func Page(items []model.Sample, offset, limit int) []model.Sample {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	if offset >= len(items) {
		return []model.Sample{}
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	return items[offset:end]
}
