package main

import (
	"ds/stack"
	"fmt"
)

func main() {
	s := stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())  //3
	fmt.Println(s.Peek()) //2
	fmt.Println(s.Length())
	fmt.Println(s.IsEmpty()) //false
	fmt.Println(s)
	s.Push("Nagarjun")
	fmt.Println(s)
	fmt.Println(s.Peek()) //Nagarjun

}
