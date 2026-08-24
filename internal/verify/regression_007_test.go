package verify

import (
	"context"
	"errors"
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/ingest"
)

func TestBug007QueueReturnsJobError(t *testing.T) {
	queue := ingest.New(1)
	defer queue.Close()
	expected := errors.New("instrument rejected sample")
	done := make(chan error, 1)
	if err := queue.Submit(context.Background(), ingest.Job{ID: "sample-1", Done: done, Run: func(context.Context) error { return expected }}); err != nil {
		t.Fatal(err)
	}
	if err := <-done; !errors.Is(err, expected) {
		t.Fatalf("completion error = %v, want %v", err, expected)
	}
}
