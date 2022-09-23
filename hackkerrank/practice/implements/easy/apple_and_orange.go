package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var countApples []int
	var countOranges []int
	for i := 0; i < len(apples); i++ {
		temp := apples[i] + a
		if temp >= s && temp <= t {
			countApples = append(countApples, int(temp))
		}
	}
	for i := 0; i < len(oranges); i++ {
		temp := oranges[i] + b
		if temp >= s && temp <= t {
			countOranges = append(countOranges, int(temp))
		}
	}
	fmt.Println(len(countApples))
	fmt.Println(len(countOranges))
}

func main() {
	countApplesAndOranges(7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4})
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
