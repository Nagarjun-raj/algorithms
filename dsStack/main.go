package main

import (
	"ds/stack"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	s := stack.New(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Data)
	err := s.Push(5)
	checkError(err)
	v, err1 := s.Pop()
	checkError(err1)
	fmt.Println(v)
	p, err4 := s.Peek()
	checkError(err4)
	fmt.Println(p)
	fmt.Println(s.Length())
	fmt.Println(s.IsEmpty())

}
