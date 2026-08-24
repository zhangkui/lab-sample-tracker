package service

import (
	"context"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"sync"
)

type BatchService struct {
	lab     *Lab
	batches map[string]model.Batch
	mu      sync.Mutex
}

func NewBatchService(l *Lab) *BatchService {
	return &BatchService{lab: l, batches: map[string]model.Batch{}}
}
func (b *BatchService) Create(id string, sampleIDs []string) model.Batch {
	// Copy the caller's slice so later mutations to the input cannot
	// leak into the stored batch. Using len==cap guarantees that any
	// subsequent append reallocates rather than writing into the
	// backing array shared with stored/returned batches.
	samples := make([]string, len(sampleIDs))
	copy(samples, sampleIDs)
	v := model.Batch{ID: id, SampleIDs: samples, State: "open"}
	b.mu.Lock()
	b.batches[id] = v
	b.mu.Unlock()
	return v
}
func (b *BatchService) Add(ctx context.Context, id, sampleID string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	v := b.batches[id]
	v.SampleIDs = append(v.SampleIDs, sampleID)
	b.batches[id] = v
	return nil
}
func (b *BatchService) Get(id string) (model.Batch, bool) { v, ok := b.batches[id]; return v, ok }
