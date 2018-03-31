package dataStructures

import "errors"

var (
	// ErrInvalidCall is generic invalid call to function error
	ErrInvalidCall = errors.New("invalid call, refer to documentation")
	// ErrStackOverFlow is stack overflow error - i.e. trying to push an element when element count is at capacity
	ErrStackOverFlow = errors.New("stack overflow")
	// ErrStackUnderFlow is stack underflow error - i.e. trying to pop an element when element count is already at 0
	ErrStackUnderFlow = errors.New("stack underflow")
	// ErrQueueOverFlow is returned when a queue overflows
	ErrQueueOverFlow = errors.New("queue overflow")
	// ErrQueueUnderFlow is returned when a queue underflows
	ErrQueueUnderFlow = errors.New("queue underflow")
)
