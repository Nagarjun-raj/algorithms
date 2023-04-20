package main

import (
	"ds/queue"
	"fmt"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	q := queue.New(4)
	q.EnQueue("Hi")
	q.EnQueue("Hello")
	q.EnQueue("Good Morning")
	q.EnQueue("Nagarjun")
	fmt.Println(q.SizeOfQueue())
	v, err := q.DeQueue()
	errorCheck(err)
	fmt.Println(v)
	f, err1 := q.Front()
	errorCheck(err1)
	fmt.Println(f)
	fmt.Println(q.IsEmpty())
}
