package verify

import (
	"sync"
	"testing"

	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
)

func TestBug001ProtocolCatalogRace(t *testing.T) {
	catalog := service.NewProtocolCatalog()
	if err := catalog.Put(model.Protocol{ID: "p", Name: "baseline", Version: 1, Steps: []model.ProtocolStep{{Name: "extract", Required: true}}}); err != nil {
		t.Fatal(err)
	}
	var group sync.WaitGroup
	group.Add(2)
	go func() {
		defer group.Done()
		for version := 2; version < 500; version++ {
			_ = catalog.Put(model.Protocol{ID: "p", Name: "updated", Version: version, Steps: []model.ProtocolStep{{Name: "extract", Required: true}}})
		}
	}()
	go func() {
		defer group.Done()
		for attempt := 0; attempt < 500; attempt++ {
			_, _ = catalog.Get("p")
		}
	}()
	group.Wait()
}
