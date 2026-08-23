package service

import (
	"context"
	"github.com/jb843051627/lab-sample-tracker/internal/ingest"
	"github.com/jb843051627/lab-sample-tracker/internal/model"
)

func (b *BatchService) Process(ctx context.Context, id string) error {
	batch, ok := b.Get(id)
	if !ok {
		return model.ErrNotFound
	}
	for _, sampleID := range batch.SampleIDs {
		done := make(chan error, 1)
		if err := b.lab.queue.Submit(ctx, ingest.Job{ID: sampleID, Done: done, Run: func(context.Context) error { _, err := b.lab.Get(sampleID); return err }}); err != nil {
			return err
		}
		if err := <-done; err != nil {
			return err
		}
	}
	return nil
}
