package verify

import (
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/service"
)

func TestBug006BatchOwnsSampleIDs(t *testing.T) {
	input := []string{"sample-1", "sample-2"}
	batches := service.NewBatchService(nil)
	batches.Create("batch-1", input)
	input[0] = "replaced"
	batch, ok := batches.Get("batch-1")
	if !ok {
		t.Fatal("batch missing")
	}
	if batch.SampleIDs[0] != "sample-1" {
		t.Fatalf("batch changed after caller mutation: %q", batch.SampleIDs[0])
	}
}
