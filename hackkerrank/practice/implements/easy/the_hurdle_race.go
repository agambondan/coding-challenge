package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})
	var result int32
	if height[len(height)-1]-k > 0 {
		result = height[len(height)-1] - k
	} else {
		result = 0
	}
	return result
}

func main() {
	var k int32 = 4
	var height = []int32{1, 6, 3, 5, 2}
	fmt.Println(hurdleRace(k, height))
}
