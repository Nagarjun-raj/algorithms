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
	l := len(s.Data)
	if l < s.Size {
		s.Data = append([]interface{}{ele}, s.Data...)
		return nil
	} else {
		return errors.New("Stack is full")
	}
}

func (s *Stack) Pop() (interface{}, error) { //removes 1st element
	l := len(s.Data)
	if l == 0 {
		return nil, errors.New("Stack is empty")
	} else {

		index := 0
		ele := (s.Data)[index]
		s.Data = (s.Data)[index+1:]
		return ele, nil
	}
}

func (s *Stack) Peek() (interface{}, error) { //returns 1st element
	l := len(s.Data)
	if l == 0 {
		return nil, errors.New("Stack is empty")
	} else {

		index := 0
		ele := (s.Data)[index]
		return ele, nil
	}
}

func (s *Stack) Length() int { //returns size of stack
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool { //tells whether stack is emplty or not.
	return len(s.Data) == 0
}
