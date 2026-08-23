package service

import (
	"context"
	"encoding/json"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"io"
)

func (l *Lab) Import(ctx context.Context, r io.Reader) error {
	var values []model.Sample
	if err := json.NewDecoder(r).Decode(&values); err != nil {
		return err
	}
	for _, v := range values {
		if err := l.Register(v); err != nil {
			return err
		}
	}
	return nil
}
