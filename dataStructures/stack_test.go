package dataStructures

import "testing"

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		// initialSize && maxCapacity
		parms         []int
		expectedError error
	}{
		{
			"NewStack, maxCapacity: 5",
			[]int{5},
			nil,
		},
		{
			"NewStack, initialSize: 1, maxCapacity: 5",
			[]int{1, 5},
			nil,
		},
		{
			"NewStack, initialSize: 6, maxCapacity: 5",
			[]int{6, 5},
			ErrInvalidCall,
		},
		{
			"NewStack, no arguments",
			[]int{},
			ErrInvalidCall,
		},
		{
			"NewStack, 3 arguments",
			[]int{1, 2, 3},
			ErrInvalidCall,
		},
		{
			"NewStack, initialSize: -1, maxCapacity: 5",
			[]int{-1, 5},
			ErrInvalidCall,
		},
		{
			"NewStack, initialSize: 0, maxCapacity: 0",
			[]int{0, 0},
			ErrInvalidCall,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewStack(tc.parms...)
			if err != tc.expectedError {
				t.Error(err)
			}
		})
	}
}

func TestPushAndPop(t *testing.T) {
	tests := []struct {
		name string
		// initialSize && maxCapacity
		parms        []int
		data         []interface{}
		pushOneByOne bool
	}{
		{
			"NewStack, initial size: 0, maxCapacity: 10, push all together",
			[]int{0, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			false,
		},
		{
			"NewStack, initial size: 5, maxCapacity: 10, push all together",
			[]int{5, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			false,
		},
		{
			"NewStack, initial size: 0, maxCapacity: 10, push one by one",
			[]int{0, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			true,
		},
		{
			"NewStack, initial size: 10, maxCapacity: 10, push one by one",
			[]int{10, 10},
			[]interface{}{1, 2, 3, "4", "5", "6", 7, 8, 9, 10},
			true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewStack(tc.parms...)
			if err != nil {
				t.Error(err)
			}
			if tc.pushOneByOne {
				for _, v := range tc.data {
					s.push(v)
				}
			} else {
				s.Push(tc.data...)
			}
			if s.GetElementCount() != len(tc.data) {
				t.Errorf("element count: %v, number of items: %v\n", s.GetElementCount(), len(tc.data))
			}
			for ix := len(tc.data) - 1; ix >= 0; ix-- {
				p, err := s.Pop()
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

func TestStackOverFlow(t *testing.T) {
	s, err := NewStack(5, 5)
	if err != nil {
		t.Error(err)
	}
	err = s.Push(1, 2, "3", "4", 5)
	if err != nil {
		t.Error(err)
	}
	err = s.Push("should fail now")
	if err != ErrStackOverFlow {
		t.Errorf("err is %v, expected error was ErrStackOverFlow\n", err)
	}
	if s.GetElementCount() != 5 {
		t.Errorf("expected element count was 5, actual element count is %v\n", s.GetElementCount())
	}
}

func TestStackUnderFlow(t *testing.T) {
	s, err := NewStack(5, 5)
	if err != nil {
		t.Error(err)
	}
	a := []interface{}{1, 2, "3", "4", 5}
	err = s.Push(a...)
	if err != nil {
		t.Error(err)
	}
	for ix := 4; ix >= 0; ix-- {
		p, err := s.Pop()
		if err != nil {
			t.Errorf("err: %v, at element index: %v\n", err, ix)
		}
		if p != a[ix] {
			t.Errorf("expected pop result was: %v, actual pop result is: %v\n", a[ix], p)
		}
	}
	_, err = s.Pop()
	if err != ErrStackUnderFlow {
		t.Errorf("err is %v, expected error was ErrStackUnderFlow\n", err)
	}
	if s.GetElementCount() != 0 {
		t.Errorf("expected element count was 0, actual element count is %v\n", s.GetElementCount())
	}
}
