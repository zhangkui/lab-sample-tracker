package model

import "errors"

var (
	ErrNotFound          = errors.New("sample not found")
	ErrInvalid           = errors.New("invalid sample")
	ErrConflict          = errors.New("sample version conflict")
	ErrInvalidTransition = errors.New("invalid status transition")
)

func ValidStatus(s SampleStatus) bool {
	switch s {
	case StatusReceived, StatusProcessing, StatusStored, StatusRejected:
		return true
	}
	return false
}
func CanTransition(from, to SampleStatus) bool {
	if from == to {
		return true
	}
	switch from {
	case StatusReceived:
		return to == StatusProcessing || to == StatusRejected
	case StatusProcessing:
		return to == StatusStored || to == StatusRejected
	case StatusStored, StatusRejected:
		return false
	}
	return false
}
