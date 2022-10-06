package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'cutTheSticks' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	var temp = arr
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	var result []int32

	for i := 0; i < len(arr); i++ {
		var count, tempValue int32
		if arr[i] != 0 {
			tempValue = arr[i]
		} else {
			continue
		}
		for j := 0; j < len(temp); j++ {
			if temp[j]-tempValue >= 0 {
				temp[j] = temp[j] - tempValue
				count += 1
			}
		}
		result = append(result, count)
		count = 0
	}

	return result
}

func main() {
	fmt.Println(cutTheSticks([]int32{5, 4, 4, 2, 2, 8}))
}
