package main

import "fmt"

/*
 * Complete the 'saveThePrisoner' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. INTEGER s
 */

func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	var result int32
	if (m+s-1)%n == 0 {
		result = n
	} else {
		result = (m + s - 1) % n
	}
	return result
}

func main() {
	fmt.Println(saveThePrisoner(7, 19, 2))
	fmt.Println(saveThePrisoner(3, 394274638, 3))
	fmt.Println(saveThePrisoner(7, 615562705, 2))

}
