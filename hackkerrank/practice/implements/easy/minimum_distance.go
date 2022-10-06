package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'minimumDistances' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func minimumDistances(a []int32) int32 {
	// Write your code here
	var result []int32
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				result = append(result, int32(j-i))
			}
		}
	}
	if len(result) < 1 {
		result = append(result, -1)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result[0]
}

func main() {
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
}
