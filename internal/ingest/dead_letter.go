package ingest

import (
	"sync"
	"time"
)

type DeadLetter struct {
	JobID    string
	Reason   string
	Payload  []byte
	FailedAt time.Time
	Attempts int
}
type DeadLetters struct {
	mu    sync.RWMutex
	items []DeadLetter
}

func (d *DeadLetters) Add(v DeadLetter) {
	d.mu.Lock()
	defer d.mu.Unlock()
	v.Payload = append([]byte(nil), v.Payload...)
	if v.FailedAt.IsZero() {
		v.FailedAt = time.Now().UTC()
	}
	d.items = append(d.items, v)
}
func (d *DeadLetters) List() []DeadLetter {
	d.mu.RLock()
	defer d.mu.RUnlock()
	out := make([]DeadLetter, len(d.items))
	copy(out, d.items)
	return out
}
func (d *DeadLetters) Remove(jobID string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i, v := range d.items {
		if v.JobID == jobID {
			d.items = append(d.items[:i], d.items[i+1:]...)
			return true
		}
	}
	return false
}
