package main

import "fmt"

//func countingValleys(steps int32, path string) int32 {
//	// Write your code here
//	var countValley, countDown, countUp int32
//	var tempStep = string(path[0])
//	if tempStep == "D" {
//		countDown = 1
//	} else {
//		countUp = 1
//	}
//	for i := 1; i < int(steps); i++ {
//		s := string(path[i])
//		//fmt.Println(tempStep, s)
//		if tempStep == s {
//			if tempStep == "D" {
//				countDown += 2
//			} else {
//				countUp += 2
//			}
//		}
//		if countDown >= 2 && countUp >= 2 {
//			fmt.Println("COUNT UP", countUp)
//			countValley = countUp / 2
//		}
//		tempStep = string(path[i])
//		//fmt.Println(tempStep, s)
//	}
//	fmt.Println(countValley)
//
//	return countValley
//}
func countingValleys(steps int32, path string) int32 {
	var counts, currentPos int32
	seaLevel := true

	for i := int32(0); i < steps; i++ {
		if path[i] == 'U' {
			currentPos++
		} else {
			currentPos--
		}

		if seaLevel && currentPos == -1 {
			counts++
		}

		seaLevel = currentPos == 0
	}

	return counts
}

func main() {
	//var path = "UDDDUDUU"
	var path = "DDUUDDUDUUUD"
	fmt.Println(countingValleys(int32(len(path)), path))
	var path1 = "DUDDDUUDUU"
	fmt.Println(countingValleys(int32(len(path1)), path1))
	var path2 = "DUDDDUUDUU"
	fmt.Println(countingValleys(int32(len(path2)), path2))
}
