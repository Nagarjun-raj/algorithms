package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	a1 := l.PushBack("is an")
	a2 := l.PushFront("to learn")
	l.InsertAfter("interesting", a2)
	l.InsertBefore("language", a1)
	a3 := l.PushFront("clearly")
	l.MoveAfter(a1, a2)
	l.MoveToBack(a2)
	l.PushBack("Golang")
	temp := l.Back()
	l.MoveToFront(temp)
	l.Remove(a3)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//Output : Golang is an interesting language to learn
	fmt.Println("\nSize of the list:", l.Len())
}
