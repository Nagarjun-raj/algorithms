package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()     //creating a list
	l.PushBack("Hello") //adding elements at last
	x := l.PushBack("Nagarjun")
	l.PushFront(1)                   //adding element at front
	b := l.Back()                    //returns address of last element
	f := l.Front()                   //returns address of front element
	l.InsertAfter("Good morning", x) //adds element after x address element
	l.MoveBefore(b, f)               //moves "Good morning" before "Hello"
	size := l.Len()                  //size of list
	fmt.Println(size)
	for e := l.Front(); e != nil; e = e.Next() { //iterating over list e.Next() is the next element address in list
		fmt.Println(e.Value) //printing the value for address just e [Nagarjun 1 Hello Good morning]
	}
	fmt.Println()
	//in reverse
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
