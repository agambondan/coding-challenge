package main

import "fmt"

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */
// https://medium.com/@mrunankmistry52/non-divisible-subset-problem-comprehensive-explanation-c878a752f057
// nonDivisibleSubset function
func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	var result int32
	var count = make(map[int32]int32)
	for i, value := range s {
		fmt.Print(value, i, " | ")
		remainder := value % k
		count[remainder] += 1
	}
	if count[0] != 0 {
		result = 1
	}

	if k%2 == 0 {
		result += 1
	}
	for i := 1; i < int(k/2+1); i++ {
		if i != int(k)-i {
			if count[int32(i)] > count[k-int32(i)] {
				result += count[int32(i)]
			} else {
				result += count[k-int32(i)]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(nonDivisibleSubset(3, []int32{1, 7, 2, 4}))
	fmt.Println(nonDivisibleSubset(3, []int32{1, 2, 3, 4, 5, 6}))
	fmt.Println(nonDivisibleSubset(6, []int32{12, 6, 1, 9, 13, 15, 10, 21, 14, 32, 5, 8, 23, 19}))
	fmt.Println(nonDivisibleSubset(5, []int32{2, 7, 12, 17, 22}))
	fmt.Println(nonDivisibleSubset(7, []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}))
}

//func nonDivisibleSubset(k int32, s []int32) int32 {
//	// Write your code here
//	var key = make(map[int32]int32)
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			if (s[i]+s[j])%k != 0 {
//				//fmt.Print(s[i]+s[j], " | ")
//				key[s[i]] = s[i] + s[j]
//				//key[s[j]] = s[i] + s[j]
//			} else if (s[i]+s[j])%k == 0 {
//				//delete(key, s[i])
//			}
//		}
//	}
//	fmt.Println(key)
//	return int32(len(key))
//}
