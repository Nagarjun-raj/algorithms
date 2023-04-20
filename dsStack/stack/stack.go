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

func (s *Stack) Push(ele interface{}) error { //adds elements on top
	l := s.Length()
	if l < s.Size {
		s.Data = append([]interface{}{ele}, s.Data...)
		return nil
	}
	return errors.New("Stack is full")
}

func (s *Stack) Pop() (interface{}, error) { //removes 1st element
	l := s.Length()
	if l == 0 {
		return nil, errors.New("Stack is empty")
	}
	ele := (s.Data)[0]
	s.Data = (s.Data)[1:]
	return ele, nil
}

func (s *Stack) Peek() (interface{}, error) { //returns 1st element
	l := s.Length()
	if l == 0 {
		return nil, errors.New("Stack is empty")
	}
	ele := (s.Data)[0]
	return ele, nil
}

func (s *Stack) Length() int { //returns size of stack
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool { //tells whether stack is emplty or not.
	return s.Length() == 0
}
