package verify

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/zhangkui/lab-sample-tracker/internal/ingest"
)

func TestBug010SubmitHonorsCancellationWhenFull(t *testing.T) {
	queue := ingest.New(1)
	release := make(chan struct{})
	defer func() { close(release); queue.Close() }()
	if err := queue.Submit(context.Background(), ingest.Job{ID: "blocking", Run: func(context.Context) error { <-release; return nil }}); err != nil {
		t.Fatal(err)
	}
	if err := queue.Submit(context.Background(), ingest.Job{ID: "buffered", Run: func(context.Context) error { return nil }}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result := make(chan error, 1)
	go func() {
		result <- queue.Submit(ctx, ingest.Job{ID: "cancelled", Run: func(context.Context) error { return nil }})
	}()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("submit error = %v, want context cancellation", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("submit did not honor cancelled context")
	}
}
