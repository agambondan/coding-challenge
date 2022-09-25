package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	var result, currentPosition int32
	lenC := len(c)
	if lenC < 3 {
		return 1
	}
	for i := 0; i < lenC-2; i++ {
		if currentPosition+2 < int32(lenC) {
			if c[currentPosition+2] == 1 {
				currentPosition += 1
			} else {
				currentPosition += 2
			}
			result += 1
		} else if currentPosition+2 == int32(lenC) {
			result += 1
			break
		}
	}
	return result
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0}))
}
