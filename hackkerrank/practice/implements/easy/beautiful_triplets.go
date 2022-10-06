package main

import "fmt"

/*
 * Complete the 'beautifulTriplets' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func contains(arr []int32, element int32) bool {
	var result = false
	for _, x := range arr {
		if x == element {
			result = true
			break
		}
	}
	return result
}

func beautifulTriplets(d int32, arr []int32) int32 {
	// Write your code here
	var result int32
	for i := 0; i < len(arr); i++ {
		if contains(arr, arr[i]+d) && contains(arr, arr[i]+d*2) {
			result += 1
		}
		//for j := i + 1; j < len(arr); j++ {
		//	for k := j + 1; k < len(arr); k++ {
		//		if arr[j]-arr[i] == d && arr[k]-arr[j] == d {
		//			result += 1
		//		}
		//	}
		//}
	}
	return result
}

func main() {
	fmt.Println(beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10}))
}
