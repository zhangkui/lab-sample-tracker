package verify

import (
	"sync"
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/metrics"
)

func TestBug002MetricsSnapshotRace(t *testing.T) {
	registry := metrics.New()
	var group sync.WaitGroup
	group.Add(2)
	go func() {
		defer group.Done()
		for index := 0; index < 1000; index++ {
			registry.Add("received", 1)
		}
	}()
	go func() {
		defer group.Done()
		for index := 0; index < 1000; index++ {
			_ = registry.Snapshot()
		}
	}()
	group.Wait()
}
