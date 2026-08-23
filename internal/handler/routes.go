package handler

import (
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"net/http"
)

func Routes(l *service.Lab) http.Handler {
	m := http.NewServeMux()
	m.Handle("/samples", ListHandler(l))
	return m
}
