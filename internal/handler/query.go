package handler

import (
	"github.com/zhangkui/lab-sample-tracker/internal/model"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"net/http"
	"strconv"
)

func ListHandler(l *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		status := model.SampleStatus(r.URL.Query().Get("status"))
		write(w, 200, l.List(status, offset, limit))
	}
}
