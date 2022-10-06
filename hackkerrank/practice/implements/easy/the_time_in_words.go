package main

import "fmt"

/*
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func timeInWords(h int32, m int32) string {
	// Write your code here
	var result string
	var nums = []string{"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen", "twenty", "twenty one",
		"twenty two", "twenty three", "twenty four",
		"twenty five", "twenty six", "twenty seven",
		"twenty eight", "twenty nine",
	}
	if m == 0 {
		result = fmt.Sprintf("%s o' clock", nums[h])
	} else if m == 1 {
		result = fmt.Sprintf("one minute past %s", nums[h])
	} else if m == 59 {
		result = fmt.Sprintf("one minute to %s", nums[(h%12)+1])
	} else if m == 15 {
		result = fmt.Sprintf("quarter past %s", nums[h])
	} else if m == 30 {
		result = fmt.Sprintf("half past %s", nums[h])
	} else if m == 45 {
		result = fmt.Sprintf("quarter to %s", nums[(h%12)+1])
	} else if m <= 30 {
		result = fmt.Sprintf("%s minutes past %s", nums[m], nums[h])
	} else if m > 30 {
		result = fmt.Sprintf("%s minutes to %s", nums[60-m],
			nums[(h%12)+1])
	}
	return result
}

func main() {
	fmt.Println(timeInWords(7, 15))
	fmt.Println(timeInWords(3, 00))
}
