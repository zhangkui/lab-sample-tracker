package verify

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"github.com/zhangkui/lab-sample-tracker/internal/store"
)

func TestBug004CreateBatchInitializesIndex(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "lab.json"))
	if err != nil {
		t.Fatal(err)
	}
	lab := service.NewLab(repo)
	defer lab.Close()
	if _, err := lab.NewIntake().CreateBatch("project-1", "batch-1", time.Now()); err != nil {
		t.Fatal(err)
	}
}
