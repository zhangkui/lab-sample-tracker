package ingest

import (
	"context"
	"fmt"
	"time"
)

type BatchJob2 struct {
	ID        string
	CreatedAt time.Time
	Attempts  int
	Payload   []byte
}

func (j BatchJob2) Validate() error {
	if j.ID == "" {
		return fmt.Errorf("job id is required")
	}
	if len(j.Payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	return nil
}

func (j BatchJob2) Run(ctx context.Context, fn func([]byte) error) error {
	if err := j.Validate(); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return fn(append([]byte(nil), j.Payload...))
}
