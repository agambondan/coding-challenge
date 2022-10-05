package main

import (
	"fmt"
)

/*
 * Complete the 'libraryFine' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d1
 *  2. INTEGER m1
 *  3. INTEGER y1
 *  4. INTEGER d2
 *  5. INTEGER m2
 *  6. INTEGER y2
 */

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	// Write your code here
	var result int32
	if m1 == m2 && y1 == y2 && d1 > d2 {
		result = (d1 - d2) * 15
	} else if y1 == y2 && m1 > m2 {
		result = (m1 - m2) * 500
	} else if y1 > y2 {
		result = 10000
	} else {
		result = 0
	}
	return result
}

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015))
	fmt.Println(libraryFine(6, 6, 2015, 9, 6, 2016))
	fmt.Println(libraryFine(2, 7, 1014, 1, 1, 1015), "a")
	fmt.Println(libraryFine(2, 7, 1014, 1, 1, 1014))
	fmt.Println(libraryFine(5, 5, 2014, 23, 2, 2014))
	fmt.Println(libraryFine(31, 5, 2014, 1, 5, 2014))
	fmt.Println(libraryFine(1, 1, 2015, 31, 12, 2014))
	fmt.Println(libraryFine(31, 8, 2004, 20, 1, 2004))
	fmt.Println(libraryFine(2, 5, 2015, 30, 5, 2015))
	fmt.Println(libraryFine(28, 2, 2015, 15, 4, 2015))
	fmt.Println(libraryFine(2, 6, 2014, 5, 7, 2014))
}

//func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
//	// Write your code here
//	var result int32
//	month1 := time.Month(int(m1))
//	month2 := time.Month(int(m2))
//	date1 := time.Date(int(y1), month1, int(d1), 0, 0, 0, 0, time.Local)
//	date2 := time.Date(int(y2), month2, int(d2), 0, 0, 0, 0, time.Local)
//	diff := date2.Sub(date1)
//	if diff.Hours()*-1 > 0 {
//		if 24*31 > diff.Hours()*-1 {
//			result = int32(15 * (diff.Hours() / 24 * -1))
//		} else if 24*30*12 > diff.Hours()*-1 {
//			round := math.Round(diff.Hours() / 24 / 30 * -1)
//			round += 0.5
//			fmt.Println(round < diff.Hours()/24/30*-1, round, diff.Hours()/24/30*-1)
//			if round < diff.Hours()/24/30*-1 {
//				round += 1
//				fmt.Println("MASOK COK", round)
//			}
//			round -= 0.5
//			result = int32(500 * round)
//		} else {
//			result = 10000
//		}
//	}
//	return result
//}
