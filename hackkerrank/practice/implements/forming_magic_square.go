package main

import (
	"fmt"
	"sort"
)

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	// a row must 45/3 = 15
	// 3x3 matrix value must 45 which is |1,9| mean 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45
	var cost int
	var tempCost int
	var countCost []int
	var tempCountCost []int
	var t = [][][]int32{
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	}
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			for k := 0; k < len(t[i][j]); k++ {
				tempCost = int(t[i][j][k] - s[j][k])
				if tempCost < 0 {
					tempCost = tempCost * -1
				}
				tempCountCost = append(tempCountCost, tempCost)
			}
		}
		for j := 0; j < len(tempCountCost); j++ {
			cost = cost + tempCountCost[j]
		}
		countCost = append(countCost, cost)
		tempCountCost = []int{}
		cost = 0
	}
	sort.Ints(countCost)
	return int32(countCost[0])
}

func main() {
	var s = [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}
	fmt.Println(formingMagicSquare(s))
}

//func formingMagicSquare(s [][]int32) int32 {
//	// Write your code here
//	// a row must 45/3 = 15
//	// 3x3 matrix value must 45 which is |1,9| mean 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 = 45
//	var cost int32
//	var minCost int32
//	var count = make([][]int, 3, 8)
//	var countTemp int
//	var t = [][][]int32{
//		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
//		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
//		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
//		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
//		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
//		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
//		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
//		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
//	}
//	for i := 0; i < len(t); i++ {
//		for j := 0; j < len(t[i]); j++ {
//			count[j] = make([]int, 8)
//			for k := 0; k < len(t[i][j]); k++ {
//				//fmt.Println(count[i], "WKWK")
//				fmt.Print(t[i][j][k], s[j][k], ", ")
//				cost = t[i][j][k] - s[j][k]
//				//fmt.Print(cost, " COK ")
//				if cost < 0 {
//					cost = cost * -1
//				}
//				//fmt.Print(cost, " HMM ")
//				count[j][] = int(cost)
//				fmt.Println(count[j], "WKWK")
//			}
//			sort.Ints(count[j])
//			fmt.Println()
//		}
//	}
//	fmt.Println(count)
//	return minCost
//}
//
//func main() {
//	var s = [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}
//	fmt.Println(formingMagicSquare(s))
//}
