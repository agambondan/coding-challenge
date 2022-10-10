package main

import (
	"fmt"
	"math"
)

// flatlandSpaceStations the maximum distance any city is from a space station
// https://oscarzhou1989.medium.com/hackerrank-flatland-space-stations-solution-dded103bb9e3
// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int32) int32 {
	var m, i, max = int32(len(c)), int32(0), int32(0)
	if n == m {
		return 0
	}
	for i = 0; i < n; i++ {
		var min = n // max distance between any of two cities won't be longer than n
		// step 1: Find the minimum distance of all cities to their nearest space station
		for _, station := range c {
			d := math.Abs(float64(i - station))
			if min > int32(d) {
				min = int32(d)
			}
		}
		// step 2: Find the maximum distance from the step 1 result
		if max < min {
			max = min
		}
	}
	return max
}

func main() {
	fmt.Println(flatlandSpaceStations(5, []int32{0, 4}))
	fmt.Println(flatlandSpaceStations(20, []int32{13, 1, 11, 10, 6}))
	fmt.Println(flatlandSpaceStations(100, []int32{93, 41, 91, 61, 30, 6, 25, 90, 97}))
}

//// Complete the flatlandSpaceStations function below.
//func flatlandSpaceStations(n int32, c []int32) int32 {
//	var result []int32
//	var mapResult = make(map[int32]struct{})
//	if int(n) == len(c) {
//		return 0
//	}
//	sort.Slice(c, func(i, j int) bool {
//		return c[i] < c[j]
//	})
//	for i := 0; i < len(c)-1; i++ {
//		mapResult[c[i+1]-c[i]+1] = struct{}{}
//		if i == 0 {
//			mapResult[c[i]] = struct{}{}
//		} else {
//			mapResult[c[i]-c[i-1]+1] = struct{}{}
//		}
//	}
//	for key := range mapResult {
//		result = append(result, key)
//	}
//	sort.Slice(result, func(i, j int) bool {
//		return result[i] > result[j]
//	})
//	fmt.Println(result, mapResult)
//	return result[0]
//}
