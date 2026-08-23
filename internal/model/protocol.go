package model

import (
	"errors"
	"sort"
	"strings"
)

type ProtocolStep struct {
	Name       string
	Required   bool
	MinSeconds int
	MaxSeconds int
}
type Protocol struct {
	ID         string
	Name       string
	Version    int
	Steps      []ProtocolStep
	Deprecated bool
}

func (p Protocol) Validate() error {
	if strings.TrimSpace(p.ID) == "" || strings.TrimSpace(p.Name) == "" || p.Version < 1 {
		return errors.New("invalid protocol")
	}
	if len(p.Steps) == 0 {
		return errors.New("protocol has no steps")
	}
	for _, s := range p.Steps {
		if s.Name == "" || s.MinSeconds < 0 || s.MaxSeconds < s.MinSeconds {
			return errors.New("invalid protocol step")
		}
	}
	return nil
}
func (p Protocol) OrderedSteps() []ProtocolStep {
	out := append([]ProtocolStep(nil), p.Steps...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
func (p Protocol) RequiredStepCount() int {
	n := 0
	for _, s := range p.Steps {
		if s.Required {
			n++
		}
	}
	return n
}
