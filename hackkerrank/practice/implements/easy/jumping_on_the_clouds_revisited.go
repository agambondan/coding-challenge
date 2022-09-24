package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32, k int32) int32 {
	initLengthC := len(c)
	var temp = c
	if initLengthC%int(k) != 0 {
		for i := 0; i < int(k)-(initLengthC%int(k)); i++ {
			c = append(c, temp...)
		}
	}
	var result int32 = 100
	for i := 0; i < len(c)/int(k); i++ {
		if c[((i+1)*int(k))%len(c)] == 0 {
			result -= 1
		} else {
			result -= 1 + 2
		}
	}
	return result
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2))
	fmt.Println(jumpingOnClouds([]int32{0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0}, 3))
	fmt.Println(jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3))
}
