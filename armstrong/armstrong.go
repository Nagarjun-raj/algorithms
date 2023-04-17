package main

import (
	"fmt"
	"math"
)

func armstrongCheck(n int) bool {
	n1 := float64(n / 100)
	n2 := float64((n % 100) / 10)
	n3 := float64(n % 10)
	calculated := math.Pow(n1, 3) + math.Pow(n2, 3) + math.Pow(n3, 3)
	//fmt.Println(calculated)
	return float64(n) == calculated
}

func main() {
	num := 371
	if armstrongCheck(num) {
		fmt.Printf("%d is an Armstrong number\n", num)
		return
	}
	fmt.Printf("%d is not an Armstrong number\n", num)
}
