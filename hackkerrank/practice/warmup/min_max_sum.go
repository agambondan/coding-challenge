package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var result []int
	for i := 0; i < len(arr); i++ {
		var temp int
		for j := 0; j < len(arr); j++ {
			if i != j {
				temp += int(arr[j])
			}
		}
		result = append(result, temp)
		temp = 0
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	fmt.Println(result[0], result[len(result)-1])
}

func main() {
	var arr = []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
	var arr1 = []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr1)
	var arr2 = []int32{5, 5, 5, 5, 5}
	miniMaxSum(arr2)
}
