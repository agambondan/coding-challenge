package main

import (
	"fmt"
	"problem_solving/lib"
	"sort"
)

/*
 * Complete the 'equalizeArray' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func equalizeArray(arr []int32) int32 {
	// Write your code here
	var keyArr = make(map[int32]int32)
	var result int32
	for i := 0; i < len(arr); i++ {
		keyArr[arr[i]] += 1
	}
	array := lib.MapValueToArray(keyArr)
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	result = int32(len(arr)) - array[len(array)-1]
	return result
}

func main() {
	fmt.Println(equalizeArray([]int32{1, 2, 3, 1, 2, 3, 3, 3}))
}
