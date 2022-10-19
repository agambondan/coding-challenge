package main

import (
	"fmt"
)

/*
 * Complete the 'strangeCounter' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER t as parameter.
 */

func strangeCounter(t int64) int64 {
	// Write your code here
	var currentTime, time, value, i int64 = 3, 3, 3, 1
	for time < t {
		value *= 2
		time += value
		if i > 1 {
			currentTime = time - value
		}
		i += 1
	}
	fmt.Println(time, value, currentTime)
	if currentTime+1 == t {
		return value
	}
	return time - t + 1
}

func main() {
	fmt.Println(strangeCounter(1))
	fmt.Println(strangeCounter(2))
	fmt.Println(strangeCounter(3))
	fmt.Println(strangeCounter(4))
	fmt.Println(strangeCounter(10))
	fmt.Println(strangeCounter(12))
	fmt.Println(strangeCounter(1000))
	fmt.Println(strangeCounter(1000000000000))
}
