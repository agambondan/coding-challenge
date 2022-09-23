package main

import "fmt"

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for i := 0; i < int(n); i++ {
		// print of space
		// loop until len n - index - 1
		for j := 1; j < int(n)-i; j++ {
			fmt.Print(" ")
		}
		// print the symbol
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	staircase(6)
}
