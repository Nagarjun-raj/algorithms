/* Maria plays college basketball and wants to go pro. Each season she maintains a record of her play. She tabulates the number of times she breaks her season record for most points and least points in a game.
Points scored in the first game establish her record for the season, and she begins counting from there.
Given the scores for a season, determine the number of times Maria breaks her records for most and least points scored during the season.

Returns
int[2]: An array with the numbers of times she broke her records. Index 0 is for breaking most points records, and index 1 is for breaking least points records.
Sample Input 0
10 5 20 20 4 5 2 25 1
Sample Output
2 4

She broke her best record twice (after games 2  and 7 ) and her worst record four times (after games 1 ,4 ,6 , and 8),
 so we print 2 4 as our answer. Note that she did not break her record for best score during
 game 3, as her score during that game was not strictly greater than her best record at the time.

*/

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
