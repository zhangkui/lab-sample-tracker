package handler

import (
	"encoding/json"
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"net/http"
	"strings"
)

type Handler struct{ lab *service.Lab }

func New(l *service.Lab) http.Handler {
	h := &Handler{lab: l}
	m := http.NewServeMux()
	m.HandleFunc("/healthz", h.health)
	m.HandleFunc("/samples", h.samples)
	m.HandleFunc("/samples/", h.sample)
	m.HandleFunc("/report", h.report)
	return m
}
func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	write(w, 200, map[string]string{"status": "ok"})
}
func (h *Handler) samples(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write(w, 200, h.lab.List("", 0, 100))
		return
	}
	if r.Method != "POST" {
		write(w, 405, nil)
		return
	}
	var v model.Sample
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		write(w, 400, map[string]string{"error": err.Error()})
		return
	}
	if err := h.lab.Register(v); err != nil {
		write(w, 400, map[string]string{"error": err.Error()})
		return
	}
	write(w, 201, v)
}
func (h *Handler) sample(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/samples/")
	v, err := h.lab.Get(id)
	if err != nil {
		write(w, 404, map[string]string{"error": err.Error()})
		return
	}
	write(w, 200, v)
}
func (h *Handler) report(w http.ResponseWriter, _ *http.Request) { write(w, 200, h.lab.Report()) }
