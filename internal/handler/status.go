package handler

import (
	"context"
	"encoding/json"
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"github.com/jb843051627/lab-sample-tracker/internal/service"
	"net/http"
	"strconv"
	"strings"
)

func StatusHandler(l *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) != 3 {
			http.NotFound(w, r)
			return
		}
		version, _ := strconv.Atoi(r.URL.Query().Get("version"))
		var req struct {
			Status model.SampleStatus `json:"status"`
		}
		if json.NewDecoder(r.Body).Decode(&req) != nil {
			http.Error(w, "bad request", 400)
			return
		}
		if err := l.Transition(context.Background(), parts[1], req.Status, version); err != nil {
			http.Error(w, err.Error(), 409)
			return
		}
		w.WriteHeader(204)
	}
}
