package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	suffix := s[8:]
	hour, _ := strconv.Atoi(s[:2])
	if suffix == "PM" && hour != 12 {
		s = fmt.Sprintf("%d%s", hour+12, s[2:8])
	} else {
		s = fmt.Sprintf("%s", s[:8])
	}
	if suffix == "AM" && hour == 12 {
		s = fmt.Sprintf("0%d%s", hour-12, s[2:8])
	}
	return s
}

func main() {
	dateTime := "07:05:45PM"
	fmt.Println(timeConversion(dateTime))
	dateTime1 := "12:40:22AM"
	fmt.Println(timeConversion(dateTime1))
	dateTime2 := "12:45:54PM"
	fmt.Println(timeConversion(dateTime2))
}
