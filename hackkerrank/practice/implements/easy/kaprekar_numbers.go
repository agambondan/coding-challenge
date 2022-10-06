package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'kaprekarNumbers' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER q
 */

func kaprekarNumbers(p int32, q int32) []int {
	// Write your code here
	var result []int
	var count int
	for i := int(p); i <= int(q); i++ {
		square := strconv.Itoa(i * i)
		first := square[:len(square)/2]
		second := square[len(square)/2:]
		firstConv, _ := strconv.Atoi(first)
		secondConv, _ := strconv.Atoi(second)
		if firstConv+secondConv == i {
			result = append(result, i)
			fmt.Print(i, " ")
			count += 1
		}
	}
	if count == 0 {
		fmt.Print("INVALID RANGE")
	}
	fmt.Println()
	return result
}

func main() {
	fmt.Println(kaprekarNumbers(1, 99999))
}
