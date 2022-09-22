package main

import "fmt"

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var result int32
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result
}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
