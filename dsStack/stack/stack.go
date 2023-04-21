package stack

import "errors"

type Stack struct {
	Size int
	Data []interface{}
}

func New(size int) *Stack {
	return &Stack{
		Size: size,
		Data: make([]interface{}, 0),
	}
}

// adds elements on top
func (s *Stack) Push(ele interface{}) error {
	l := s.Length()
	if l < s.Size {
		s.Data = append([]interface{}{ele}, s.Data...)
		return nil
	}
	return errors.New("Stack is full")
}

// removes 1st element
func (s *Stack) Pop() (interface{}, error) {
	l := s.Length()
	if l == 0 {
		return nil, errors.New("Stack is empty")
	}
	ele := (s.Data)[0]
	s.Data = (s.Data)[1:]
	return ele, nil
}

// returns 1st element
func (s *Stack) Peek() (interface{}, error) {
	l := s.Length()
	if l == 0 {
		return nil, errors.New("Stack is empty")
	}
	ele := (s.Data)[0]
	return ele, nil
}

// returns size of stack
func (s *Stack) Length() int {
	return len(s.Data)
}

// tells whether stack is emplty or not.
func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}
