package main

import (
	"fmt"
)

func unique(intSlice []int32) []int32 {
	l := len(intSlice)
	keys := make(map[int32]struct{})
	var list []int32
	for i := 1; i <= l; i++ {
		val := intSlice[l-i]
		if _, value := keys[val]; !value {
			keys[val] = struct{}{}
			list = append(list, val)
		}
	}
	//for _, entry := range intSlice {
	//	if _, value := keys[entry]; !value {
	//		keys[entry] = struct{}{}
	//		list = append(list, entry)
	//	}
	//}
	//sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	return list
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	newRanked := unique(ranked)
	//var result = make([]int32, len(player))
	var result []int32
	if len(newRanked) == 1 {
		for i := 0; i < len(player); i++ {
			if player[i] > newRanked[0] {
				//result[i] = 1
				result = append(result, 1)
			} else if player[0] == newRanked[0] {
				//result[i] = 1
				result = append(result, 1)
			} else if player[0] < newRanked[0] {
				//result[i] = 2
				result = append(result, 2)
			}
		}
	} else {
		for i := 0; i < len(player); i++ {
			l := len(newRanked)
			l32 := int32(l)
			p := player[i]
			var temp int32
			for j := 1; j < l; j++ {
				if p > newRanked[j] {
					temp = 1
				} else if p > newRanked[j-1] && p < newRanked[j] {
					temp = l32 - int32(j) + 1
					break
				} else if p == newRanked[j-1] {
					temp = l32 - int32(j) + 1
					break
				} else if p < newRanked[j-1] {
					temp = l32 + 1
					break
				}
			}
			result = append(result, temp)
			//result[i] = temp
			temp = 0
		}
	}
	return result
}

// TODO: how to optimize this code
func main() {
	var ranked = []int32{100, 100, 50, 40, 40, 20, 10}
	var player = []int32{5, 25, 50, 125}
	fmt.Println(climbingLeaderboard(ranked, player))
	var ranked1 = []int32{1}
	var player1 = []int32{1, 1}
	fmt.Println(climbingLeaderboard(ranked1, player1))
}

//func contains(s []int32, e int32) bool {
//	for _, a := range s {
//		if a == e {
//			return true
//		}
//	}
//	return false
//}
//
//func remove(slice []int32, s int) []int32 {
//	return append(slice[:s], slice[s+1:]...)
//}
//
//func add(slice []int32, s int, add int32) []int32 {
//	return append(append(slice[:s], add), slice[s+1:]...)
//}
//
//func climbingLeaderboard(ranked []int32, player []int32) []int32 {
//	// Write your code here
//	for i := 0; i < len(ranked); i++ {
//		if contains(ranked[i+1:], ranked[i]) {
//			ranked = remove(ranked, i)
//			i--
//		}
//	}
//	sort.Slice(ranked, func(i, j int) bool { return ranked[i] < ranked[j] })
//	var result = make([]int32, len(player))
//	if len(ranked) == 1 {
//		for i := 0; i < len(player); i++ {
//			if player[i] > ranked[0] {
//				result[i] = 1
//			} else if player[0] == ranked[0] {
//				result[i] = 1
//			} else if player[0] < ranked[0] {
//				result[i] = 2
//			}
//		}
//	} else {
//		for i := 0; i < len(player); i++ {
//			l := len(ranked)
//			l32 := int32(l)
//			p := player[i]
//			var temp int32
//			for j := 1; j < l; j++ {
//				if p > ranked[j] {
//					temp = 1
//				} else if p > ranked[j-1] && p < ranked[j] {
//					temp = l32 - int32(j) + 1
//					break
//				} else if p == ranked[j-1] {
//					temp = l32 - int32(j) + 1
//					break
//				} else if p < ranked[j-1] {
//					temp = l32 + 1
//					break
//				}
//			}
//			result[i] = temp
//			temp = 0
//		}
//	}
//	return result
//}

//func climbingLeaderboard(ranked []int32, player []int32) []int32 {
//	// Write your code here
//	ranked = unique(ranked)
//	lP := len(player)
//	lNR := len(ranked)
//	var result = make([]int32, lP)
//	if lNR == 1 {
//		for i := 0; i < len(player); i++ {
//			if player[i] > ranked[0] {
//				result[i] = 1
//			} else if player[0] == ranked[0] {
//				result[i] = 1
//			} else if player[0] < ranked[0] {
//				result[i] = 2
//			}
//		}
//	} else {
//		for i := 0; i < len(player); i++ {
//			l := len(ranked)
//			l32 := int32(l)
//			p := player[i]
//			fmt.Print(p, " | ")
//			var temp int32
//			for j := 0; j < l; j++ {
//				fmt.Print(ranked[l-j-1], ranked[l-i-1], ranked[j], " | ")
//				if p > ranked[l-j-1] {
//					temp = 1
//				} else if p < ranked[l-j-1] && p > ranked[l-i-1] {
//					temp = l32 - int32(j) + 1
//					break
//				} else if p == ranked[l-i-1] {
//					temp = l32 - int32(j) + 1
//					break
//				} else if p < ranked[l-i-1] {
//					temp = l32 + 1
//					break
//				}
//			}
//			result[i] = temp
//			temp = 0
//			fmt.Println()
//		}
//	}
//	return result
//}
