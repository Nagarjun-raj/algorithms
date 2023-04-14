package main

import "fmt"

func calculate(scores []int) []int {
	var high, low, hcount, lcount int
	result := []int{0, 0}
	highest := make([]int, 0)
	lowest := make([]int, 0)
	max := scores[0]
	for _, v := range scores {
		if v >= max {
			highest = append(highest, v)
			max = v
		} else {
			highest = append(highest, max)
		}
	}

	min := scores[0]
	for _, v := range scores {
		if v <= min {
			lowest = append(lowest, v)
			min = v
		} else {
			lowest = append(lowest, min)
		}
	}
	high = highest[0]
	for i := 1; i < len(highest); i++ {
		if highest[i] != high {
			hcount++
			high = highest[i]
		}
	}

	low = lowest[0]
	for i := 1; i < len(lowest); i++ {
		if lowest[i] != low {
			lcount++
			low = lowest[i]
		}
	}

	result[0], result[1] = hcount, lcount

	return result
}

func main() {
	scores := []int{10, 5, 20, 20, 4, 5, 2, 25, 1}
	record := calculate(scores)
	fmt.Println(record)

}
