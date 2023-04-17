/*
Given an array of integers, calculate the ratios of its elements that are positive, negative and zero.
Print the decimal value of each fraction on a new line with  places after the decimal.
*/
package main

import "fmt"

func plusMinus(arr []int) {
	n := float32(len(arr))
	var pcount, ncount, zcount float32
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pcount++
		} else if arr[i] < 0 {
			ncount++
		} else {
			zcount++
		}
	}
	positive_ratio := pcount / n
	fmt.Printf("Positive ratio:%.6f\n", positive_ratio)
	negative_ratio := ncount / n
	fmt.Printf("Negative ratio:%.6f\n", negative_ratio)
	zero_ratio := zcount / n
	fmt.Printf("Zero ratio:%.6f\n", zero_ratio)
}

func main() {
	arr := []int{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
