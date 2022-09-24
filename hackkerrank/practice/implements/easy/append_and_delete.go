package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	fmt.Printf("%s | %s | %d | %d | %d | ", s, t, len(s), len(t), k)
	// Write your code here
	var result string
	lenS := len(s)
	lenT := len(t)
	if strings.EqualFold(s, t) {
		if int(k) >= lenS*2 || k%2 == 0 {
			result = "Yes"
		} else {
			result = "No"
		}
		return result
	}
	var temp int
	var length int
	if lenS > lenT {
		length = lenT
	} else {
		length = lenS
	}
	for i := 0; i < length; i++ {
		if s[i] != t[i] {
			break
		}
		temp += 1
	}
	tempS := lenS - temp
	tempT := lenT - temp
	tot := tempS + tempT
	kInt := int(k)
	if (tot == kInt) || (tot < kInt && (tot-kInt)%2 == 0) || (tot+(2*temp) <= kInt) {
		result = "Yes"
	} else {
		result = "No"
	}
	return result
}

func main() {
	fmt.Println(appendAndDelete("hackerhappy", "hackerrank", 9), "Expected : Yes")
	fmt.Println(appendAndDelete("aba", "aba", 7), "Expected : Yes")
	fmt.Println(appendAndDelete("ashley", "ash", 2), "Expected : No")
	fmt.Println(appendAndDelete("y", "yu", 2), "Expected : No")
	fmt.Println(appendAndDelete("zzzzz", "zzzzzzz", 4), "Expected : Yes")
	fmt.Println(appendAndDelete("qwerasdf", "qwerbsdf", 6), "Expected : No")
	fmt.Println(appendAndDelete("asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
		"bsdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv", 100), "Expected : No")
	fmt.Println(appendAndDelete("abcd", "abcdert", 10), "Expected : No")

}

//func appendAndDelete(s string, t string, k int32) string {
//	fmt.Printf(" %s %s %d %d %d | ", s, t, len(s), len(t), k)
//	// Write your code here
//	var result string
//	lenS := len(s)
//	lenT := len(t)
//	var temp int32
//	for i := 0; i < lenS; i++ {
//		if lenT > i {
//			if s[i] != t[i] {
//				temp = int32(lenT - i)
//				break
//			} else {
//				temp = 0
//			}
//		} else {
//			temp = int32(lenS - i)
//			break
//		}
//	}
//	fmt.Printf("%d | ", temp)
//	if temp > k && (int(k) < lenS || int(k) < lenT) {
//		result = "No"
//	} else {
//		result = "Yes"
//	}
//	return result
//}

//func appendAndDelete(s string, t string, k int32) string {
//	fmt.Printf("%s | %s | %d | %d | %d | ", s, t, len(s), len(t), k)
//	// Write your code here
//	var result string
//	var temp int32
//	lenS := len(s)
//	lenT := len(t)
//	if lenS > lenT {
//		for i := 0; i < lenT; i++ {
//			if s[i] != t[i] {
//				temp = int32(lenT - i)
//				break
//			}
//		}
//	} else if lenS < lenT {
//
//	} else if lenS == lenT {
//
//	}
//	fmt.Printf("temp : %d |", temp)
//
//	return result
//}
