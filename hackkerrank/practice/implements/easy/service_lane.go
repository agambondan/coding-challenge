package main

import "fmt"

/*
 * Complete the 'serviceLane' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY cases
 */

func serviceLane(width []int32, cases [][]int32) []int32 {
	// Write your code here
	var result []int32
	for i := 0; i < len(cases); i++ {
		var temp = width[cases[i][0]]
		for j := cases[i][0] + 1; j <= cases[i][1]; j++ {
			if width[j] < temp {
				temp = width[j]
			}
		}
		result = append(result, temp)
		temp = 0
	}
	return result
}

func main() {
	fmt.Println(serviceLane([]int32{2, 3, 1, 2, 3, 2, 3, 3}, [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}))
	fmt.Println(serviceLane([]int32{1, 2, 2, 2, 1}, [][]int32{{2, 3}, {1, 4}, {2, 4}, {2, 4}, {2, 3}}))
}
