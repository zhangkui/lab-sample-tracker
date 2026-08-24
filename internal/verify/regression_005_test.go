package verify

import (
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/ingest"
)

func TestBug005ManifestOwnsFileList(t *testing.T) {
	files := []string{"first.csv", "second.csv"}
	manifest := ingest.NewManifest("manifest-1", files)
	files[0] = "replaced.csv"
	if manifest.Files[0] != "first.csv" {
		t.Fatalf("manifest changed after caller mutation: %q", manifest.Files[0])
	}
}
