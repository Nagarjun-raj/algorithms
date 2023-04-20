package main

import (
	"fmt"
	"sort"
)

func tallestCandle(arr []int) int {
	sort.Ints(arr)
	count := 0
	max := arr[len(arr)-1]
	for _, v := range arr {
		if v == max {
			count++
		}
	}
	return count

}

func main() {
	arr := make([]int, 0)
	var age int
	var height int
	fmt.Println("Enter the age of birthday boy: ")
	fmt.Scanln(&age)
	for i := 0; i < age; i++ {
		fmt.Println("Enter height unit of candle", i+1)
		fmt.Scanln(&height)
		arr = append(arr, height)
	}
	x := tallestCandle(arr)
	fmt.Println("Number of tallest candles are: ", x)

}
