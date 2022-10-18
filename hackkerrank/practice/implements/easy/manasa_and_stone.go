package main

import (
	"fmt"
)

/*
 * Complete the 'stones' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER a
 *  3. INTEGER b
 */

/*
	Big O Notation is 2 to the power of (n-1)
	Big O Notation adalah 2 pangkat (n-1)
	6 4 8
	case 4 to 8
	0, 4, 8, 12, 16, 20
	0, 4, 8, 12, 16, 24
	0, 4, 8, 12, 20, 28
	0, 4, 8, 16, 24, 32
	0, 4, 12, 20, 28, 36

	0, 4, 8, 12, 20, 24
	0, 4, 16, 24, 32, 36
	0, 4, 16, 24, 28, 32
	0, 4, 16, 20, 24, 28
	0, 4, 12, 16, 20, 24

	case 8 to 4
	0, 8, 16, 24, 32, 40
	0, 8, 16, 24, 32, 36
	0, 8, 16, 24, 28, 32
	0, 8, 16, 20, 24, 28
	0, 8, 12, 16, 20, 24

	case 4 between 8
	0, 4, 16, 24, 32, 40
	0, 4, 16, 24, 32, 36
	0, 4, 16, 24, 28, 32
	0, 4, 16, 20, 24, 28
	0, 4, 12, 16, 20, 24
*/
func recursiveSquare(n int32) int32 {
	if n == 0 {
		return 0
	}
	return recursiveSquare(n-1) + n + n - 1
}

func stones(n int32, a int32, b int32) []int32 {
	// Write your code here
	if a > b {
		a, b = b, a
	}
	var result []int32
	var initVal = a * (n - 1)
	var gap = b - a
	result = append(result, initVal)
	if gap == 0 {
		return result
	}
	for i := 1; i < int(n); i++ {
		result = append(result, initVal+(int32(i)*gap))
	}
	return result
}

func main() {
	T := [][]int32{{3, 1, 2}, {4, 10, 100}, {6, 4, 8}, {11, 3, 10}}
	for i := 0; i < len(T); i++ {
		fmt.Println(stones(T[i][0], T[i][1], T[i][2]))
	}
	fmt.Println(recursiveSquare(2))
	fmt.Println(recursiveSquare(5))
}
