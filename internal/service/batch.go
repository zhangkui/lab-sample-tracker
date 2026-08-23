package service

import (
	"context"
	"github.com/jb843051627/lab-sample-tracker/internal/model"
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
	v := model.Batch{ID: id, SampleIDs: append([]string(nil), sampleIDs...), State: "open"}
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
