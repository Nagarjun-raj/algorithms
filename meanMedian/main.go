package main

import (
	"log"
	"sort"
	"sync"
	"time"
)

func addNums(i int, w *sync.WaitGroup, ch chan int) {
	time.Sleep(time.Second)
	ch <- i
	w.Done()
}

func meanMedianCal(ch chan int, ch2 chan bool, count int, sum int) {
	var median = 0
	temp := make([]int, 0)
	for j := range ch {
		temp = append(temp, j)
		sort.Ints(temp)
		count++
		sum += j
		log.Println(j)
		mean := float64(sum) / float64(count)
		time.Sleep(time.Second)
		log.Printf("Mean=%.2f", mean)
		if count%2 != 0 {
			middle := (len(temp) - 1) / 2
			median = temp[middle]
			time.Sleep(time.Second)
			log.Println("Median=", median)
		} else {
			middle := len(temp) / 2
			median = (temp[middle] + temp[middle-1]) / 2
			time.Sleep(time.Second)
			log.Println("median=", median)

		}

	}
	ch2 <- true

}

func main() {
	var count = 0
	var sum = 0
	ch := make(chan int)
	var w sync.WaitGroup
	ch2 := make(chan bool)
	go meanMedianCal(ch, ch2, count, sum)
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	for _, v := range arr {
		w.Add(1)
		go addNums(v, &w, ch)
	}
	w.Wait()
	close(ch)
	<-ch2

}
