package queue

import "testing"

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		// initialSize && maxCapacity
		parms         []int
		expectedError error
	}{
		{
			"NewQueue, maxCapacity: 5",
			[]int{5},
			nil,
		},
		{
			"NewQueue, initialSize: 1, maxCapacity: 5",
			[]int{1, 5},
			nil,
		},
		{
			"NewQueue, initialSize: 6, maxCapacity: 5",
			[]int{6, 5},
			ErrInvalidCall,
		},
		{
			"NewQueue, no arguments",
			[]int{},
			ErrInvalidCall,
		},
		{
			"NewQueue, 3 arguments",
			[]int{1, 2, 3},
			ErrInvalidCall,
		},
		{
			"NewQueue, initialSize: -1, maxCapacity: 5",
			[]int{-1, 5},
			ErrInvalidCall,
		},
		{
			"NewQueue, initialSize: 0, maxCapacity: 0",
			[]int{0, 0},
			ErrInvalidCall,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewQueue(tc.parms...)
			if err != tc.expectedError {
				t.Error(err)
			}
		})
	}
}

func TestEnqueueAndDequeue(t *testing.T) {
	tests := []struct {
		name string
		// initialSize && maxCapacity
		parms         []int
		data          []interface{}
		queueOneByOne bool
	}{
		{
			"NewQueue, initial size: 0, maxCapacity: 10, push all together",
			[]int{0, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			false,
		},
		{
			"NewQueue, initial size: 5, maxCapacity: 10, push all together",
			[]int{5, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			false,
		},
		{
			"NewQueue, initial size: 0, maxCapacity: 10, push one by one",
			[]int{0, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			true,
		},
		{
			"NewQueue, initial size: 10, maxCapacity: 10, push one by one",
			[]int{10, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q, err := NewQueue(tc.parms...)
			if err != nil {
				t.Error(err)
			}
			if tc.queueOneByOne {
				for _, v := range tc.data {
					q.enqueue(v)
				}
			} else {
				q.Enqueue(tc.data...)
			}
			if q.GetElementCount() != len(tc.data) {
				t.Errorf("element count: %v, number of items: %v\n", q.GetElementCount(), len(tc.data))
			}
			for ix := 0; ix < len(tc.data); ix++ {
				p, err := q.Dequeue()
				if err != nil {
					t.Error(err)
				}
				if p != tc.data[ix] {
					t.Errorf("expected element: %v, element found: %v", tc.data[ix], p)
				}
			}
		})
	}
}

func TestQueueOverFlow(t *testing.T) {
	q, err := NewQueue(5, 5)
	if err != nil {
		t.Error(err)
	}
	err = q.Enqueue(1, 2, "3", "4", 5)
	if err != nil {
		t.Error(err)
	}
	err = q.Enqueue("should fail now")
	if err != ErrQueueOverFlow {
		t.Errorf("err is %v, expected error was ErrQueueOverFlow\n", err)
	}
	if q.GetElementCount() != 5 {
		t.Errorf("expected element count was 5, actual element count is %v\n", q.GetElementCount())
	}
}

func TestQueueUnderFlow(t *testing.T) {
	q, err := NewQueue(5, 5)
	if err != nil {
		t.Error(err)
	}
	a := []interface{}{1, 2, "3", "4", 5}
	err = q.Enqueue(a...)
	if err != nil {
		t.Error(err)
	}
	for ix := 0; ix <= 4; ix++ {
		p, err := q.Dequeue()
		if err != nil {
			t.Errorf("err: %v, at element index: %v\n", err, ix)
		}
		if p != a[ix] {
			t.Errorf("expected pop result was: %v, actual pop result is: %v\n", a[ix], p)
		}
	}
	_, err = q.Dequeue()
	if err != ErrQueueUnderFlow {
		t.Errorf("err is %v, expected error was ErrQueueUnderFlow\n", err)
	}
	if q.GetElementCount() != 0 {
		t.Errorf("expected element count was 0, actual element count is %v\n", q.GetElementCount())
	}
}
