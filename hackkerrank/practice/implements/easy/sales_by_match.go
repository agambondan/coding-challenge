package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var result int32
	var socks = make(map[string]int32)

	for i := 0; i < int(n); i++ {
		key := strconv.Itoa(int(ar[i]))
		_, exist := socks[key]
		if exist {
			socks[key] += 1
		} else {
			socks[key] = 1
		}
	}
	for _, v := range socks {
		if v/2 > 0 {
			result += v / 2
		}
	}
	return result
}

func main() {
	var arr = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(int32(len(arr)), arr))
}
