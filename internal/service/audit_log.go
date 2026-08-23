package service

import (
	"errors"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"sort"
	"time"
)

func (l *Lab) record(subject, action string, details map[string]string) error {
	if subject == "" || action == "" {
		return errors.New("audit subject and action are required")
	}
	return l.repo.Event(model.Event{SampleID: subject, Action: action, Actor: "workflow", At: l.clock.Now(), Details: details})
}
func (l *Lab) Timeline(subject string) []model.Event {
	events := l.repo.Events(subject)
	sort.SliceStable(events, func(i, j int) bool { return events[i].At.Before(events[j].At) })
	return events
}
func (l *Lab) TimelineSince(subject string, since time.Time) []model.Event {
	out := []model.Event{}
	for _, e := range l.Timeline(subject) {
		if !e.At.Before(since) {
			out = append(out, e)
		}
	}
	return out
}
