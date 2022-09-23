package main

import "fmt"

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	// Write your code here
	for i := 0; i < int(n); i++ {
		for j := i; j < int(n); j++ {
			if i < j && (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	var ar = []int32{1, 3, 2, 6, 1, 2}
	fmt.Println(divisibleSumPairs(int32(len(ar)), 3, ar))

}
