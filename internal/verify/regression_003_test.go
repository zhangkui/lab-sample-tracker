package verify

import (
	"path/filepath"
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"github.com/zhangkui/lab-sample-tracker/internal/store"
)

func TestBug003RegisterInitializesCache(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "lab.json"))
	if err != nil {
		t.Fatal(err)
	}
	lab := service.NewLab(repo)
	defer lab.Close()
	if err := lab.Register(model.Sample{ID: "sample-1", Subject: "subject-1"}); err != nil {
		t.Fatal(err)
	}
}
