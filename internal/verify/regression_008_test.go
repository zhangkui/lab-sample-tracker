package verify

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"github.com/zhangkui/lab-sample-tracker/internal/store"
)

func TestBug008StartRejectsBusyInstrument(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "lab.json"))
	if err != nil {
		t.Fatal(err)
	}
	lab := service.NewLab(repo)
	defer lab.Close()
	workflow := lab.NewAssayWorkflow()
	if err := workflow.RegisterProtocol(model.Protocol{ID: "pcr", Name: "PCR", Version: 1, Steps: []model.ProtocolStep{{Name: "extract", Required: true}}}); err != nil {
		t.Fatal(err)
	}
	if err := workflow.RegisterInstrument(model.Instrument{ID: "inst-1", Name: "Instrument", State: model.InstrumentOnline, Capabilities: []string{"pcr"}}); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"assay-1", "assay-2"} {
		if err := workflow.Queue(model.AssayRequest{ID: id, SampleID: id, ProtocolID: "pcr", InstrumentID: "inst-1"}); err != nil {
			t.Fatal(err)
		}
	}
	if err := workflow.Start(context.Background(), "assay-1", time.Now()); err != nil {
		t.Fatal(err)
	}
	if err := workflow.Start(context.Background(), "assay-2", time.Now()); err == nil {
		t.Fatal("busy instrument accepted a second assay")
	}
}
