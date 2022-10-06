package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	// Write your code here
	var result int64
	if len(s) == 1 && s == "a" {
		return n
	}
	if len(s) == 1 && s != "a" {
		return 0
	}
	countRepeated := n / int64(len(s))
	countA := int64(strings.Count(s, "a"))
	result = countRepeated * countA
	if n%int64(len(s)) != 0 {
		result += int64(strings.Count(s[:n%int64(len(s))], "a"))
	}
	fmt.Println(len(s), countRepeated, countRepeated*countA, result)
	return result
}

func main() {
	fmt.Println(repeatedString("a", 1000000000000))
	fmt.Println(repeatedString("aba", 10))
	fmt.Println(repeatedString("kmretasscityylpdhuwjirnqimlkcgxubxmsxpypgzxtenweirknjtasxtvxemtwxuarabssvqdnktqadhyktagjxoanknhgilnm", 736778906400))
}
