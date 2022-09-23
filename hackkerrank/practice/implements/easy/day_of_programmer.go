package main

import "fmt"

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here
	var date string
	var programmerDay = 256
	var monthOneToEight = []int{31, 31, 30, 31, 30, 31, 31}
	if year > 1918 {
		if year%400 == 0 {
			monthOneToEight = append(monthOneToEight, 29)
		} else if year%100 == 0 {
			monthOneToEight = append(monthOneToEight, 28)
		} else if year%4 == 0 {
			monthOneToEight = append(monthOneToEight, 29)
		} else {
			monthOneToEight = append(monthOneToEight, 28)
		}
	} else {
		if year%4 == 0 {
			monthOneToEight = append(monthOneToEight, 29)
		} else {
			monthOneToEight = append(monthOneToEight, 28)
		}
	}
	var dayYear int
	for _, value := range monthOneToEight {
		dayYear += value
	}
	date = fmt.Sprintf("%d.09.%d", programmerDay-dayYear, year)

	if year == 1918 {
		date = fmt.Sprintf("26.09.%d", year)
	}

	return date
}

func main() {
	fmt.Println(dayOfProgrammer(2016))
	fmt.Println(dayOfProgrammer(2017))
	fmt.Println(dayOfProgrammer(2012))
	fmt.Println(dayOfProgrammer(2100))
	fmt.Println(dayOfProgrammer(1918))
	fmt.Println(dayOfProgrammer(1800))
	fmt.Println(dayOfProgrammer(1700))
	fmt.Println(dayOfProgrammer(1917))
}
