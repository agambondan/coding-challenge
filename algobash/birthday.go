package main

import (
	"fmt"
)

func birthday(s []int32, d int32, m int32) int32 {
	var sum int32 // for storing the total no of ways
	var length = len(s)
	for i := 0; i < length; i++ {
		var sp int // for storing the sum of the m digits
		/*
		   if sum of current position and m is greater than size of s then there will be overflow
		   so need to check just return the sum
		*/
		if int32(i)+m > int32(length) {
			fmt.Println(i, "i+m lebih besar dari panjang array index =", length-1)
			return sum
		}
		for j := 0; j < int(m); j++ {
			fmt.Printf("i = %d + j = %d\n", i, j)
			fmt.Println(s[i+j])
			sp = sp + int(s[i+j])
		}
		if int32(sp) == d {
			sum++
		}
	}
	return sum
}

func main() {
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 3))
	//fmt.Println(birthday([]int32{1, 1, 1, 1, 1, 1}, 3, 3))
}
