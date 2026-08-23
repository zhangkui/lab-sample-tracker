package handler

import (
	"encoding/json"
	"github.com/jb843051627/lab-sample-tracker/internal/model"
	"net/http"
)

func EncodeSample(w http.ResponseWriter, v model.Sample) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
func DecodeSample(r *http.Request) (model.Sample, error) {
	var v model.Sample
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}
