package queue

import "errors"

var (
	// ErrInvalidCall is generic invalid call to function error
	ErrInvalidCall = errors.New("invalid call, refer to documentation")
	// ErrQueueOverFlow is returned when a queue overflows
	ErrQueueOverFlow = errors.New("queue overflow")
	// ErrQueueUnderFlow is returned when a queue underflows
	ErrQueueUnderFlow = errors.New("queue underflow")
)
