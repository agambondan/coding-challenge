package main

import "fmt"

// SolutionSwap is your solution code.
func SolutionSwap(arr []int) int {
	var count = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			for j := arr[i]; j <= i+1; j++ {
				fmt.Println(j)
				var temp = 0
				temp = arr[arr[i]-1]
				arr[arr[i]-1] = arr[i]
				arr[i] = temp
				count += 1
			}
		}
	}

	return count
}
func main() {
	var arr = []int{4, 3, 1, 2}
	fmt.Println(SolutionSwap(arr))
}
