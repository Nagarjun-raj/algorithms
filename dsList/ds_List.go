package main

import (
	"ds/nagarjun"
	"log"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	l := nagarjun.New(3)
	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(3)
	v, err := l.Remove(0)
	errorCheck(err)
	log.Println(v)
	log.Println(l)
	k, err1 := l.InsertAfter(4, 1)
	errorCheck(err1)
	log.Println(l)
	log.Println(k)
	log.Println(l)
	n, err2 := l.Remove(1)
	log.Println(n)
	errorCheck(err2)
	log.Println(l)

}
