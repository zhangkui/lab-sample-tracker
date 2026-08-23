package model

import (
	"errors"
	"sort"
	"time"
)

type BatchState string

const (
	BatchOpen     BatchState = "open"
	BatchSealed   BatchState = "sealed"
	BatchInReview BatchState = "in_review"
	BatchReleased BatchState = "released"
)

type BatchItem struct {
	SampleID  string
	Position  int
	AddedAt   time.Time
	RemovedAt *time.Time
}

type BatchRecord struct {
	ID         string
	Project    string
	State      BatchState
	Items      []BatchItem
	CreatedAt  time.Time
	SealedAt   *time.Time
	ReleasedAt *time.Time
}

func (b BatchRecord) ActiveItems() []BatchItem {
	items := make([]BatchItem, 0, len(b.Items))
	for _, item := range b.Items {
		if item.RemovedAt == nil {
			items = append(items, item)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Position < items[j].Position })
	return items
}

func (b BatchRecord) Contains(sampleID string) bool {
	for _, item := range b.Items {
		if item.SampleID == sampleID && item.RemovedAt == nil {
			return true
		}
	}
	return false
}

func (b *BatchRecord) Add(sampleID string, at time.Time) error {
	if b.State != BatchOpen {
		return errors.New("batch is not open")
	}
	if sampleID == "" || b.Contains(sampleID) {
		return errors.New("sample is already in batch")
	}
	b.Items = append(b.Items, BatchItem{SampleID: sampleID, Position: len(b.Items) + 1, AddedAt: at})
	return nil
}

func (b *BatchRecord) Seal(at time.Time) error {
	if b.State != BatchOpen || len(b.ActiveItems()) == 0 {
		return errors.New("batch cannot be sealed")
	}
	b.State = BatchSealed
	b.SealedAt = &at
	return nil
}

func (b *BatchRecord) StartReview() error {
	if b.State != BatchSealed {
		return errors.New("batch is not sealed")
	}
	b.State = BatchInReview
	return nil
}

func (b *BatchRecord) Release(at time.Time) error {
	if b.State != BatchInReview {
		return errors.New("batch is not under review")
	}
	b.State = BatchReleased
	b.ReleasedAt = &at
	return nil
}
