package main

import "fmt"

/*
 * Complete the 'utopianTree' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func utopianTree(n int32) int32 {
	// Write your code here
	var result int32 = 1
	for i := 1; i <= int(n); i++ {
		if i%2 == 0 {
			result += 1
		} else {
			result = result * 2
		}
	}
	return result
}

func main() {
	var arr = []int32{4, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(utopianTree(arr[i]))
	}
}
