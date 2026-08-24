package verify

import (
	"context"
	"errors"
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/ingest"
)

func TestBug009RetryPreservesCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	attemptError := errors.New("temporary instrument failure")
	if err := ingest.Retry(ctx, 1, func() error { return attemptError }); !errors.Is(err, context.Canceled) {
		t.Fatalf("retry error = %v, want context cancellation", err)
	}
}
