package main

import "fmt"

/*
 * Complete the 'chocolateFeast' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c
 *  3. INTEGER m
 */

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	var result, temp int32
	bars := n / c
	result = bars
	temp = bars
	for i := 0; i < int(bars); i++ {
		fmt.Printf("before : %d ", temp)
		if temp/m >= 1 {
			result += temp / m
			temp = temp%m + temp/m
			fmt.Printf("temp : %d, m : %d\n", temp, m)
		} else {
			break
		}
	}
	fmt.Print("| Result | : ")
	return result
}

func main() {
	fmt.Println(chocolateFeast(10, 2, 5))
	fmt.Println(chocolateFeast(12, 4, 4))
	fmt.Println(chocolateFeast(6, 2, 2))
	fmt.Println(chocolateFeast(16809, 123, 11668))
	fmt.Println(chocolateFeast(20373, 18211, 10188))
}
