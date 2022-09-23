package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var d1Sum, d2Sum int32
	for x := range arr {
		for y := range arr {
			if x == y {
				d1Sum += arr[x][y]
			}
			fmt.Println(x == len(arr)-y-1, x, len(arr), y)
			if x == len(arr)-y-1 {
				d2Sum += arr[x][y]
			}
		}
	}

	return int32(math.Abs(float64(d1Sum) - float64(d2Sum)))
}

func main() {
	var arr2Dimension = [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	fmt.Println(diagonalDifference(arr2Dimension))
}
