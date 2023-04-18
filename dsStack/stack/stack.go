package stack

type Stack []interface{}

func (s *Stack) Push(ele interface{}) { //adds elements on top
	*s = append([]interface{}{ele}, *s...)
}

func (s *Stack) Pop() interface{} { //removes 1st element
	index := 0
	ele := (*s)[index]
	*s = (*s)[index+1:]
	return ele
}

func (s *Stack) Peek() interface{} { //returns 1st element
	index := 0
	ele := (*s)[index]
	return ele
}

func (s *Stack) Length() int { //returns size of stack
	return len(*s)
}

func (s *Stack) IsEmpty() bool { //tells whether stack is emplty or not.
	return *s == nil
}
