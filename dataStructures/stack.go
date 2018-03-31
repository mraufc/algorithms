package dataStructures

import (
	"errors"
)

var (
	// ErrInvalidCall is generic invalid call to function error
	ErrInvalidCall = errors.New("invalid call, refer to documentation")
	// ErrStackOverFlow is stack overflow error - i.e. trying to push an element when element count is at capacity
	ErrStackOverFlow = errors.New("stack overflow")
	// ErrStackUnderFlow is stack underflow error - i.e. trying to pop an element when element count is already at 0
	ErrStackUnderFlow = errors.New("stack underflow")
)

// Stack is the stack struct.
type Stack struct {
	maxCapacity  int
	elementCount int
	elements     []interface{}
	increment    int
}

// NewStack retruns a new stack.
// Usage: NewStack(initialSize, maxCapacity) or NewStack(maxCapacity)
// initialSize should be non-negative and maxCapacity should be positive, initialSize should be less than or equal to maxCapacity
func NewStack(args ...int) (*Stack, error) {
	var i, m int
	switch len(args) {
	case 1:
		m = args[0]
	case 2:
		i = args[0]
		m = args[1]
	default:
		return nil, ErrInvalidCall
	}
	if i < 0 || m < 1 || i > m {
		return nil, ErrInvalidCall
	}

	s := &Stack{
		maxCapacity:  m,
		elementCount: 0,
		elements:     make([]interface{}, i),
		increment:    1,
	}
	return s, nil
}
func (s *Stack) push(a interface{}) error {
	if s.elementCount >= s.maxCapacity {
		return ErrStackOverFlow
	}
	if s.elementCount < len(s.elements) {
		s.elements[s.elementCount] = a
		s.elementCount++
		return nil
	}
	n := len(s.elements) + s.increment
	if n > s.maxCapacity {
		n = s.maxCapacity
	}
	s.elements = append(s.elements, make([]interface{}, n-len(s.elements))...)
	s.elements[s.elementCount] = a
	s.elementCount++
	s.increment *= 2
	return nil
}

// Push one or more elements to stack. First element is pushed first.
func (s *Stack) Push(args ...interface{}) error {
	for _, arg := range args {
		if err := s.push(arg); err != nil {
			return err
		}
	}
	return nil
}

// Pop removed and returns the latest element added to a stack
func (s *Stack) Pop() (interface{}, error) {
	if s.elementCount <= 0 {
		return nil, ErrStackUnderFlow
	}
	s.elementCount--
	return s.elements[s.elementCount], nil
}

func (s *Stack) GetElementCount() int {
	return s.elementCount
}
