package main

import "fmt"

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var result int64
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result
}

func main() {
	fmt.Println(aVeryBigSum([]int64{1, 2, 3}))
}
