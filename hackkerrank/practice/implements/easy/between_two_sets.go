package main

import "fmt"

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	n := len(a)
	m := len(b)
	var c int32
	for i := 1; i <= 100; i++ {
		factor := true
		for j := 0; j < n; j++ {
			if i%int(a[j]) != 0 {
				factor = false
				break
			}
		}
		if factor {
			for j := 0; j < m; j++ {
				if int(b[j])%i != 0 {
					factor = false
					break
				}
			}
		}
		if factor {
			c++
		}
	}
	return c
}

func main() {
	var a = []int32{2, 4}
	var b = []int32{16, 32, 96}
	fmt.Println(getTotalX(a, b))
}
