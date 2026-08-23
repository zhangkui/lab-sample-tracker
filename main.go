package main

import (
	"github.com/zhangkui/lab-sample-tracker/internal/handler"
	"github.com/zhangkui/lab-sample-tracker/internal/service"
	"github.com/zhangkui/lab-sample-tracker/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("LAB_SAMPLE_TRACKER_DB")
	if path == "" {
		path = "data/lab-sample-tracker.json"
	}
	repo, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	app := service.NewLab(repo)
	defer app.Close()
	addr := os.Getenv("LAB_SAMPLE_TRACKER_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("lab-sample-tracker listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler.New(app)))
}
