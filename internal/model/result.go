package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Name  string
	Value float64
	Unit  string
	Lower *float64
	Upper *float64
}

type InstrumentResult struct {
	ID           string
	AssayID      string
	InstrumentID string
	RunID        string
	Measurements []Measurement
	RawPayload   string
	ReceivedAt   time.Time
	Checksum     string
}

func (r *InstrumentResult) Seal() {
	sum := sha256.Sum256([]byte(r.RawPayload + r.RunID + r.InstrumentID))
	r.Checksum = hex.EncodeToString(sum[:])
}

func (r InstrumentResult) VerifySeal() bool {
	sum := sha256.Sum256([]byte(r.RawPayload + r.RunID + r.InstrumentID))
	return r.Checksum == hex.EncodeToString(sum[:])
}

func (r InstrumentResult) Validate() error {
	if r.ID == "" || r.AssayID == "" || r.InstrumentID == "" || r.RunID == "" {
		return errors.New("result identity is incomplete")
	}
	if !r.VerifySeal() {
		return errors.New("result checksum mismatch")
	}
	if len(r.Measurements) == 0 {
		return errors.New("result has no measurements")
	}
	for _, m := range r.Measurements {
		if strings.TrimSpace(m.Name) == "" || strings.TrimSpace(m.Unit) == "" {
			return errors.New("measurement is incomplete")
		}
		if m.Lower != nil && m.Upper != nil && *m.Lower > *m.Upper {
			return fmt.Errorf("invalid range for %s", m.Name)
		}
	}
	return nil
}

func (r InstrumentResult) Summary() string {
	parts := make([]string, 0, len(r.Measurements))
	for _, m := range r.Measurements {
		parts = append(parts, m.Name+"="+strconv.FormatFloat(m.Value, 'f', -1, 64)+m.Unit)
	}
	return strings.Join(parts, ", ")
}
