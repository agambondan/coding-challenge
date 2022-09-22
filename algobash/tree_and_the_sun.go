package main

import "fmt"

// Solution is your solution code.
func solution(arr []int, n int) int {
	// Your code starts here.
	result := 1
	temp := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > temp || arr[i] == temp {
			result += 1
			temp = arr[i]
		}
	}
	return result
}

func main() {
	var arr = []int{7, 4, 8, 2, 9}
	var arr1 = []int{1, 4, 4, 2, 9}
	fmt.Println(solution(arr, len(arr)))
	fmt.Println(solution(arr1, len(arr1)))
}
