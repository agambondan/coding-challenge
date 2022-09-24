package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'findDigits' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func findDigits(n int32) int32 {
	// Write your code here
	var temp []int32
	convertToString := strconv.Itoa(int(n))
	for i := 0; i < len(convertToString); i++ {
		convertToInt, _ := strconv.Atoi(string(convertToString[i]))
		temp = append(temp, int32(convertToInt))
	}
	var result int32
	for i := 0; i < len(temp); i++ {
		if temp[i] == 0 {
			continue
		}
		if n%temp[i] == 0 {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(findDigits(12))
	fmt.Println(findDigits(1012))
}
