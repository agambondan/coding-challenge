package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] < candles[j]
	})
	var result int32
	for i := 0; i < len(candles); i++ {
		if candles[len(candles)-1] == candles[i] {
			result += 1
		}
	}
	return result
}

func main() {
	var arr = []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))
}
