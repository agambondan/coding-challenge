package main

import "fmt"

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
