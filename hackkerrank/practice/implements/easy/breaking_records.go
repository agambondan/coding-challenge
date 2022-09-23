package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var result = make([]int32, 2)
	var highest, lowest = scores[0], scores[0]
	// Write your code here
	for i := 1; i < len(scores); i++ {
		if scores[i] > highest {
			highest = scores[i]
			result[0] += 1
		} else if scores[i] < lowest {
			lowest = scores[i]
			result[1] += 1
		}
	}
	return result
}

func main() {
	var scores = []int32{10, 10, 20, 20, 4, 5, 2, 3, 4, 1, 4, 4, 2, 5, 2, 1, 25, 25, 30, 25}
	fmt.Println(breakingRecords(scores))
}
