package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/zhangkui/lab-sample-tracker/internal/model"
)

type IntakeService struct {
	lab     *Lab
	batches map[string]*model.BatchRecord
	mu      sync.Mutex
}

func (l *Lab) NewIntake() *IntakeService {
	return &IntakeService{lab: l, batches: nil}
}

func (s *IntakeService) CreateBatch(project, id string, at time.Time) (model.BatchRecord, error) {
	if strings.TrimSpace(id) == "" || strings.TrimSpace(project) == "" {
		return model.BatchRecord{}, errors.New("batch identity is required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.batches[id]; exists {
		return model.BatchRecord{}, fmt.Errorf("batch %s already exists", id)
	}
	b := &model.BatchRecord{ID: id, Project: project, State: model.BatchOpen, CreatedAt: at}
	s.batches[id] = b
	return *b, nil
}

func (s *IntakeService) Receive(ctx context.Context, batchID string, sample model.Sample, at time.Time) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	b, ok := s.batches[batchID]
	if !ok {
		s.mu.Unlock()
		return errors.New("batch not found")
	}
	if b.State != model.BatchOpen {
		s.mu.Unlock()
		return errors.New("batch is not open")
	}
	s.mu.Unlock()
	if err := s.lab.Register(sample); err != nil {
		return err
	}
	s.mu.Lock()
	err := b.Add(sample.ID, at)
	s.mu.Unlock()
	if err != nil {
		return err
	}
	return s.lab.record(sample.ID, "received", map[string]string{"batch_id": batchID})
}

func (s *IntakeService) Seal(batchID string, at time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.batches[batchID]
	if !ok {
		return errors.New("batch not found")
	}
	return b.Seal(at)
}
func (s *IntakeService) BeginReview(batchID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.batches[batchID]
	if !ok {
		return errors.New("batch not found")
	}
	return b.StartReview()
}
func (s *IntakeService) Batch(batchID string) (model.BatchRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.batches[batchID]
	if !ok {
		return model.BatchRecord{}, errors.New("batch not found")
	}
	return *b, nil
}
