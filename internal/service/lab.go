package service

import (
	"context"
	"fmt"
	"github.com/jb843051627/lab-sample-tracker/internal/clock"
	"github.com/jb843051627/lab-sample-tracker/internal/ingest"
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"github.com/jb843051627/lab-sample-tracker/internal/store"
	"sync"
)

type Lab struct {
	repo    *store.Store
	clock   clock.Clock
	queue   *ingest.Queue
	mu      sync.RWMutex
	cache   map[string]model.Sample
	workers sync.WaitGroup
}

func NewLab(repo *store.Store) *Lab {
	l := &Lab{repo: repo, clock: clock.System{}, queue: ingest.New(32), cache: map[string]model.Sample{}}
	return l
}
func (l *Lab) Close() { l.queue.Close(); l.workers.Wait() }
func (l *Lab) Register(v model.Sample) error {
	if err := validate(v); err != nil {
		return err
	}
	now := l.clock.Now()
	v.UpdatedAt = now
	if v.CollectedAt.IsZero() {
		v.CollectedAt = now
	}
	if v.Status == "" {
		v.Status = model.StatusReceived
	}
	v.Version = 1
	if err := l.repo.Save(v); err != nil {
		return err
	}
	l.mu.Lock()
	l.cache[v.ID] = v
	l.mu.Unlock()
	return l.repo.Event(model.Event{SampleID: v.ID, Action: "registered", Actor: "system", At: now})
}
func (l *Lab) Get(id string) (model.Sample, error) {
	l.mu.RLock()
	v, ok := l.cache[id]
	l.mu.RUnlock()
	if ok {
		return v, nil
	}
	v, ok = l.repo.Load(id)
	if !ok {
		return model.Sample{}, model.ErrNotFound
	}
	l.mu.Lock()
	l.cache[id] = v
	l.mu.Unlock()
	return v, nil
}
func (l *Lab) Transition(ctx context.Context, id string, to model.SampleStatus, version int) error {
	return l.transition(ctx, id, to, version)
}
func (l *Lab) List(status model.SampleStatus, offset, limit int) []model.Sample {
	items := store.FilterStatus(l.repo.List(), status)
	return store.Page(store.SortByUpdate(items), offset, limit)
}
func (l *Lab) Events(id string) []model.Event { return l.repo.Events(id) }
func validate(v model.Sample) error {
	if v.ID == "" || v.Subject == "" {
		return fmt.Errorf("%w", model.ErrInvalid)
	}
	if v.Status != "" && !model.ValidStatus(v.Status) {
		return fmt.Errorf("%w", model.ErrInvalid)
	}
	return nil
}
