package main

import (
	"fmt"
	"strings"
)

// SolutionPalindrome is your solution code.
func SolutionPalindrome(str string) bool {
	// Your code starts here.
	//str = strings.Replace(strings.Trim(strings.ToLower(str), " "), "?", "", -1)
	//split := strings.Split(str, " ")
	//for i := 0; i < len(str)/2; i++ {
	//	j := len(str) - 1 - i
	//	fmt.Println(string(str[j]), string(str[i]))
	//	if str[i] != str[j] {
	//		return false
	//	}
	//}
	//return true
	str = strings.Replace(strings.Trim(strings.ToLower(str), " "), "?", "", -1)
	split := strings.Split(str, " ")
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)
	fmt.Println(str)
	for i := range split[0] {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(SolutionPalindrome("abba"))
	fmt.Println(SolutionPalindrome("test"))
	fmt.Println(SolutionPalindrome("Was it a car or a cat I saw?"))
}
