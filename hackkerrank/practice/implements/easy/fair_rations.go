package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'fairRations' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY B as parameter.
 */
// https://akshay-ravindran.medium.com/solving-fair-rations-in-o-n-52847b835dda
func fairRations(B []int32) string {
	// Write your code here
	var sum int32
	for i := 0; i < len(B); i++ {
		sum += B[i]
	}
	var result int
	if sum%2 == 1 {
		return "NO"
	} else {
		for i := 0; i < len(B); i++ {
			if B[i]%2 != 0 {
				B[i] = B[i] + 1
				B[i+1] = B[i+1] + 1
				result += 2
			}
		}
	}
	return strconv.Itoa(result)
	//if len(B) < 3 {
	//	return "NO"
	//}
	//var result = 0
	//for i := 0; i < len(B)-1; i++ {
	//	if B[i]%2 == 0 {
	//		continue
	//	} else {
	//		result += 1
	//		B[i] += 1
	//		B[i+1] += 1
	//	}
	//}
	//fmt.Println(result)
	//for i := 0; i < len(B); i++ {
	//	if B[i]%2 != 0 {
	//		result -= 1
	//		break
	//	}
	//}
	//if result > -1 {
	//	result = result * 2
	//}
	//return strconv.Itoa(result)
}

func main() {
	fmt.Println(fairRations([]int32{2, 3, 4, 5, 6}))
}
