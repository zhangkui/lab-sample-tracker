package model

import (
	"errors"
	"strings"
	"time"
)

type InstrumentState string

const (
	InstrumentOnline      InstrumentState = "online"
	InstrumentMaintenance InstrumentState = "maintenance"
	InstrumentOffline     InstrumentState = "offline"
)

type Instrument struct {
	ID            string
	Name          string
	State         InstrumentState
	LastHeartbeat time.Time
	Capabilities  []string
	ActiveRunID   string
}

func (i Instrument) CanRun(protocol string) bool {
	if i.State != InstrumentOnline || strings.TrimSpace(protocol) == "" {
		return false
	}
	for _, capability := range i.Capabilities {
		if capability == protocol {
			return true
		}
	}
	return false
}

func (i *Instrument) Heartbeat(at time.Time) {
	i.LastHeartbeat = at
	if i.State == InstrumentOffline {
		i.State = InstrumentOnline
	}
}
func (i *Instrument) Begin(runID string) error {
	if i.ActiveRunID != "" || i.State != InstrumentOnline {
		return errors.New("instrument is busy")
	}
	i.ActiveRunID = runID
	return nil
}
func (i *Instrument) End(runID string) error {
	if i.ActiveRunID != runID {
		return errors.New("run does not own instrument")
	}
	i.ActiveRunID = ""
	return nil
}
