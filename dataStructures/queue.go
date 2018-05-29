package datastructures

// Queue is the queue implementation struct. This queue implementation is not thread-safe.
type Queue struct {
	maxCapacity  int
	elementCount int
	elements     []interface{}
	increment    int
}

// NewQueue retruns a new queue.
// Usage: NewQueue(initialSize, maxCapacity) or NewQueue(maxCapacity)
// initialSize should be non-negative and maxCapacity should be positive, initialSize should be less than or equal to maxCapacity
func NewQueue(args ...int) (*Queue, error) {
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
	q := &Queue{
		maxCapacity:  m,
		elementCount: 0,
		elements:     make([]interface{}, i),
		increment:    1,
	}
	return q, nil
}
func (q *Queue) enqueue(a interface{}) error {
	if q.elementCount >= q.maxCapacity {
		return ErrQueueOverFlow
	}
	if q.elementCount < len(q.elements) {
		q.elements[q.elementCount] = a
		q.elementCount++
		return nil
	}
	n := len(q.elements) + q.increment
	doubleIncr := true
	if n > q.maxCapacity {
		n = q.maxCapacity
		doubleIncr = false
	}
	q.elements = append(q.elements, make([]interface{}, n-len(q.elements))...)
	q.elements[q.elementCount] = a
	q.elementCount++
	if doubleIncr {
		q.increment *= 2
	}
	return nil
}

// Enqueue adds one or more elements to the tail of the queue
func (q *Queue) Enqueue(args ...interface{}) error {
	for _, arg := range args {
		if err := q.enqueue(arg); err != nil {
			return err
		}
	}
	return nil
}

// Dequeue returns the element at the head of the queue.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.elementCount <= 0 {
		return nil, ErrQueueUnderFlow
	}
	r := q.elements[0]
	q.elements = q.elements[1:]
	q.elementCount--
	// there is no solid reason for me to reset increment here, however "I felt like it"
	q.increment = 1
	return r, nil
}

// GetElementCount returns the number of elements in the queue
func (q *Queue) GetElementCount() int {
	return q.elementCount
}
