package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/zhangkui/lab-sample-tracker/internal/model"
)

type ReleaseService struct {
	intake     *IntakeService
	workflow   *AssayWorkflow
	reconciler *Reconciler
}

func NewRelease(i *IntakeService, w *AssayWorkflow, r *Reconciler) *ReleaseService {
	return &ReleaseService{intake: i, workflow: w, reconciler: r}
}

func (r *ReleaseService) Release(batchID string, at time.Time) error {
	return r.release(batchID, "system", at, false)
}

func (r *ReleaseService) ForceRelease(batchID, actor string, at time.Time) error {
	if actor == "" {
		return errors.New("actor is required")
	}
	return r.release(batchID, actor, at, true)
}

func (r *ReleaseService) release(batchID, actor string, at time.Time, forced bool) error {
	b, err := r.intake.Batch(batchID)
	if err != nil {
		return err
	}
	if b.State != model.BatchInReview {
		return errors.New("batch is not under review")
	}
	assays := r.workflow.AssaysForBatch(batchID)
	if len(assays) == 0 {
		return errors.New("batch has no assays")
	}
	for _, assay := range assays {
		if assay.State != model.AssayApproved {
			return fmt.Errorf("assay %s is not approved", assay.ID)
		}
	}
	if open := r.reconciler.Open(batchID); len(open) > 0 && !forced {
		return fmt.Errorf("batch has %d open exceptions", len(open))
	}

	r.intake.mu.Lock()
	defer r.intake.mu.Unlock()
	current, ok := r.intake.batches[batchID]
	if !ok || current.State != model.BatchInReview {
		return errors.New("batch changed during release")
	}
	if err := current.Release(at); err != nil {
		return err
	}
	return r.workflow.lab.record(batchID, "batch_released", map[string]string{"actor": actor, "forced": fmt.Sprint(forced)})
}
