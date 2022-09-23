package main

import "fmt"

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	l := float64(len(arr))
	var positive, negative, zeros float64
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive += 1
		} else if arr[i] < 0 {
			negative += 1
		} else {
			zeros += 1
		}
	}
	fmt.Println(positive / l)
	fmt.Println(negative / l)
	fmt.Println(zeros / l)
}

func main() {
	var arr = []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
