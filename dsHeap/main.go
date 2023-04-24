package main

import (
	"ds/heap"
	"log"
)

func main() {
	minHeap := heap.New()
	minHeap.Push(2)
	minHeap.Push(1)
	minHeap.Push(5)
	minHeap.Push(4)
	minHeap.Push(3)
	log.Println(minHeap)
	v := minHeap.Pop()
	log.Println(v)
	log.Println(minHeap)

}
