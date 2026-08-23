package handler

import (
	"fmt"
	"net/http"
	"strings"
)

type Payload3 struct {
	Name     string
	Value    string
	Optional bool
}

func (p Payload3) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if !p.Optional && strings.TrimSpace(p.Value) == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

func (p Payload3) Header() string {
	return p.Name + ":" + p.Value
}

func WritePayload3(w http.ResponseWriter, p Payload3) {
	w.Header().Set("X-Lab-Field", p.Name)
	_, _ = w.Write([]byte(p.Value))
}
