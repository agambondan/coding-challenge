package main

import (
	"fmt"
	"problem_solving/lib"
	"time"
)

func unique(intSlice []int32) []int32 {
	l := len(intSlice)
	keys := make(map[int32]struct{})
	for i := 1; i <= l; i++ {
		val := intSlice[l-i]
		if _, value := keys[val]; !value {
			keys[val] = struct{}{}
		}
	}
	return lib.MapKeyToArray(keys)
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	newRanked := unique(ranked)
	var result []int32
	if len(newRanked) == 1 {
		for i := 0; i < len(player); i++ {
			if player[i] > newRanked[0] {
				result = append(result, 1)
			} else if player[0] == newRanked[0] {
				result = append(result, 1)
			} else if player[0] < newRanked[0] {
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
			temp = 0
		}
	}
	return result
}

// TODO: how to optimize this code
func main() {
	now := time.Now()

	var ranked = []int32{100, 100, 50, 40, 40, 20, 10}
	var player = []int32{5, 25, 50, 125}
	fmt.Println(climbingLeaderboard(ranked, player))
	fmt.Println(time.Since(now) * 1000)
	now = time.Now()
	var ranked1 = []int32{1}
	var player1 = []int32{1, 1}
	fmt.Println(climbingLeaderboard(ranked1, player1))
	fmt.Println(time.Since(now) * 1000)

	now = time.Now()
	var ranked2 = []int32{100, 100, 50, 40, 40, 20, 10}
	var player2 = []int32{5, 25, 50, 125}
	fmt.Println(climbingLeaderboard1(ranked2, player2), "MAP To array")
	fmt.Println(time.Since(now) * 1000)
	now = time.Now()
	var ranked3 = []int32{1}
	var player3 = []int32{1, 1}
	fmt.Println(climbingLeaderboard1(ranked3, player3), "MAP To array")
	fmt.Println(time.Since(now) * 1000)
}

func unique1(intSlice []int32) []int32 {
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
	return list
}

func climbingLeaderboard1(ranked []int32, player []int32) []int32 {
	// Write your code here
	newRanked := unique1(ranked)
	var result []int32
	if len(newRanked) == 1 {
		for i := 0; i < len(player); i++ {
			if player[i] > newRanked[0] {
				result = append(result, 1)
			} else if player[0] == newRanked[0] {
				result = append(result, 1)
			} else if player[0] < newRanked[0] {
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
			temp = 0
		}
	}
	return result
}
