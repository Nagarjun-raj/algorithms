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
	
	max := scores[0]
	min := scores[0]
	hCount := 0
	lCount := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			hCount++
		}
		if scores[i] < min {
			min = scores[i]
			lCount++
		}
	}

	return []int{hCount, lCount}
}

func main() {
	scores := []int{10, 5, 20, 20, 4, 5, 2, 25, 1}
	record := calculate(scores)
	fmt.Println(record)

}
