package service

import (
	"errors"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"sync"
	"time"
)

type InstrumentMonitor struct {
	mu      sync.Mutex
	items   map[string]model.Instrument
	timeout time.Duration
}

func NewInstrumentMonitor(timeout time.Duration) *InstrumentMonitor {
	return &InstrumentMonitor{items: map[string]model.Instrument{}, timeout: timeout}
}
func (m *InstrumentMonitor) Observe(i model.Instrument) {
	m.mu.Lock()
	defer m.mu.Unlock()
	old := m.items[i.ID]
	old.ID = i.ID
	old.Name = i.Name
	old.Capabilities = i.Capabilities
	old.State = i.State
	old.LastHeartbeat = i.LastHeartbeat
	old.ActiveRunID = i.ActiveRunID
	m.items[i.ID] = old
}
func (m *InstrumentMonitor) Sweep(now time.Time) []model.Instrument {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := []model.Instrument{}
	for id, i := range m.items {
		if now.Sub(i.LastHeartbeat) > m.timeout && i.State == model.InstrumentOnline {
			i.State = model.InstrumentOffline
			m.items[id] = i
			out = append(out, i)
		}
	}
	return out
}
func (m *InstrumentMonitor) Claim(id, run string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	i, ok := m.items[id]
	if !ok {
		return errors.New("instrument not found")
	}
	if err := i.Begin(run); err != nil {
		return err
	}
	m.items[id] = i
	return nil
}
