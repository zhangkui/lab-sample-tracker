package service

import (
	"context"
	"fmt"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/validation"
)

func (l *Lab) transition(ctx context.Context, id string, to model.SampleStatus, version int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	v, err := l.Get(id)
	if err != nil {
		return err
	}
	if v.Version != version {
		return model.ErrConflict
	}
	if !model.CanTransition(v.Status, to) {
		return model.ErrInvalidTransition
	}
	v.Status = to
	v.Version++
	v.UpdatedAt = l.clock.Now()
	if err = l.repo.Save(v); err != nil {
		return err
	}
	l.mu.Lock()
	l.cache[id] = v
	l.mu.Unlock()
	return l.repo.Event(model.Event{SampleID: id, Action: "status:" + string(to), Actor: "operator", At: v.UpdatedAt, Details: map[string]string{"checksum": fmt.Sprint(validation.Checksum([]byte(id)))}})
}
