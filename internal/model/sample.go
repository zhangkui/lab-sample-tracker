package model

import "time"

type SampleStatus string

const (
	StatusReceived   SampleStatus = "received"
	StatusProcessing SampleStatus = "processing"
	StatusStored     SampleStatus = "stored"
	StatusRejected   SampleStatus = "rejected"
)

type Sample struct {
	ID, Subject, SpecimenType, Location string
	Status                              SampleStatus
	CollectedAt, UpdatedAt              time.Time
	Version                             int
	Metadata                            map[string]string
}
type Event struct {
	ID                      int64
	SampleID, Action, Actor string
	At                      time.Time
	Details                 map[string]string
}
type Batch struct {
	ID                   string
	SampleIDs            []string
	CreatedAt, UpdatedAt time.Time
	State                string
}
