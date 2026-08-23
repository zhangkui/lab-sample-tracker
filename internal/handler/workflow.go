package handler

import (
	"encoding/json"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"net/http"
	"strings"
	"time"
)

type WorkflowHandler struct {
	intake   *service.IntakeService
	workflow *service.AssayWorkflow
}

func NewWorkflowHandler(i *service.IntakeService, w *service.AssayWorkflow) *WorkflowHandler {
	return &WorkflowHandler{intake: i, workflow: w}
}
func (h *WorkflowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.NotFound(w, r)
		return
	}
	switch parts[0] {
	case "batches":
		h.batch(w, r, parts[1])
	case "assays":
		h.assay(w, r, parts[1])
	default:
		http.NotFound(w, r)
	}
}
func (h *WorkflowHandler) batch(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method == http.MethodGet {
		b, err := h.intake.Batch(id)
		if err != nil {
			errorJSON(w, 404, "not_found", err.Error())
			return
		}
		write(w, 200, b)
		return
	}
	if r.Method != http.MethodPost {
		errorJSON(w, 405, "method_not_allowed", "method not allowed")
		return
	}
	var req struct {
		Sample model.Sample `json:"sample"`
	}
	if json.NewDecoder(r.Body).Decode(&req) != nil {
		errorJSON(w, 400, "bad_request", "invalid body")
		return
	}
	if err := h.intake.Receive(r.Context(), id, req.Sample, time.Now().UTC()); err != nil {
		errorJSON(w, 409, "receive_failed", err.Error())
		return
	}
	write(w, 202, map[string]string{"sample_id": req.Sample.ID})
}
func (h *WorkflowHandler) assay(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method == http.MethodGet {
		a, err := h.workflow.Assay(id)
		if err != nil {
			errorJSON(w, 404, "not_found", err.Error())
			return
		}
		write(w, 200, a)
		return
	}
	errorJSON(w, 405, "method_not_allowed", "use the assay workflow endpoints")
}
