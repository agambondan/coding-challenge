package main

import "fmt"

/*
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func reverseNumber(num int32) int32 {
	var res int32
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var result int32
	for a := i; a <= j; a++ {
		difference := a - reverseNumber(a)
		if difference < 0 {
			difference *= -1
		}
		if difference%k == 0 {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
