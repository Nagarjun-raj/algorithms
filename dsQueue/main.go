package main

import (
	"ds/queue"
	"fmt"
	"log"
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
	log.Println(q.SizeOfQueue())
	v, err := q.DeQueue()
	errorCheck(err)
	log.Println(v)
	f, err1 := q.Front()
	errorCheck(err1)
	log.Println(f)
	log.Println(q.IsEmpty())
}
