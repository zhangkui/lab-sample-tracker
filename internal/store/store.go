package store

import (
	"encoding/json"
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Store struct {
	mu      sync.RWMutex
	path    string
	samples map[string]model.Sample
	events  []model.Event
	next    int64
}

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	s := &Store{path: path, samples: map[string]model.Sample{}, next: 1}
	if b, e := os.ReadFile(path); e == nil {
		var d disk
		if json.Unmarshal(b, &d) == nil {
			s.samples = d.Samples
			s.events = d.Events
			s.next = d.Next
		}
	}
	return s, nil
}

type disk struct {
	Samples map[string]model.Sample `json:"samples"`
	Events  []model.Event           `json:"events"`
	Next    int64                   `json:"next"`
}

func (s *Store) persistLocked() error {
	b, e := json.MarshalIndent(disk{s.samples, s.events, s.next}, "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(s.path, b, 0644)
}
func (s *Store) Close() error { s.mu.Lock(); defer s.mu.Unlock(); return s.persistLocked() }
func (s *Store) Save(v model.Sample) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.samples[v.ID] = v
	return s.persistLocked()
}
func (s *Store) Load(id string) (model.Sample, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.samples[id]
	return v, ok
}
func (s *Store) List() []model.Sample {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]model.Sample, 0, len(s.samples))
	for _, v := range s.samples {
		out = append(out, v)
	}
	return out
}
func (s *Store) Event(e model.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.ID = s.next
	s.next++
	if e.At.IsZero() {
		e.At = time.Now().UTC()
	}
	s.events = append(s.events, e)
	return s.persistLocked()
}
func (s *Store) Events(id string) []model.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []model.Event{}
	for _, e := range s.events {
		if e.SampleID == id {
			out = append(out, e)
		}
	}
	return out
}
