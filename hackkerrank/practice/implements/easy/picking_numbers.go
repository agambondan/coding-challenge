package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// Write your code here
	var frequency [101]int
	var result []int
	for i := 0; i < len(a); i++ {
		index := a[i]
		fmt.Println(index)
		frequency[index] += 1
	}
	fmt.Println(frequency)

	for i := 1; i <= 100; i++ {
		result = append(result, frequency[i]+frequency[i-1])
	}
	sort.Ints(result)

	return int32(result[len(result)-1])
}

func main() {
	var arr = []int32{1, 2, 2, 3, 1, 2}
	fmt.Println(pickingNumbers(arr))
}
