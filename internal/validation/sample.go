package validation

import (
	"fmt"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"strings"
)

func Sample(v model.Sample) error {
	if strings.TrimSpace(v.ID) == "" {
		return fmt.Errorf("%w: id", model.ErrInvalid)
	}
	if strings.TrimSpace(v.Subject) == "" {
		return fmt.Errorf("%w: subject", model.ErrInvalid)
	}
	if !model.ValidStatus(v.Status) {
		return fmt.Errorf("%w: status", model.ErrInvalid)
	}
	return nil
}
