package main

import (
	"fmt"
	"strconv"
)

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var result int32
	var socks = make(map[string]int32)

	for i := 0; i < len(ar); i++ {
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
