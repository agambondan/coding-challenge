package main

import (
	"fmt"
	"sort"
	"time"
)

func contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}

func add(slice []int32, s int, add int32) []int32 {
	return append(append(slice[:s], add), slice[s+1:]...)
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	start := time.Now()
	//fmt.Println(start.String())

	//keys := make(map[int]bool)
	//var ranked []int
	//for _, entry := range ranked {
	//	e := int(entry)
	//	if _, value := keys[e]; !value {
	//		keys[e] = true
	//		ranked = append(ranked, e)
	//	}
	//}
	//sort.Ints(ranked)
	for i := 0; i < len(ranked); i++ {
		if contains(ranked[i+1:], ranked[i]) {
			ranked = remove(ranked, i)
			i--
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
	sort.Slice(ranked, func(i, j int) bool { return ranked[i] < ranked[j] })
	var result = make([]int32, len(player))
	if len(ranked) == 1 {
		for i := 0; i < len(player); i++ {
			if player[i] > ranked[0] {
				result[i] = 1
			} else if player[0] == ranked[0] {
				result[i] = 1
			} else if player[0] < ranked[0] {
				result[i] = 2
			}
		}
	} else {
		for i := 0; i < len(player); i++ {
			l := len(ranked)
			l32 := int32(l)
			p := player[i]
			var temp int32
			for j := 1; j < l; j++ {
				if p > ranked[j] {
					temp = 1
				} else if p > ranked[j-1] && p < ranked[j] {
					temp = l32 - int32(j) + 1
					break
				} else if p == ranked[j-1] {
					temp = l32 - int32(j) + 1
					break
				} else if p < ranked[j-1] {
					temp = l32 + 1
					//ranked = add(ranked, i, p)
					//ranked = append(ranked, p)
					//sort.Slice(ranked, func(i, j int) bool { return ranked[i] < ranked[j] })
					break
				}
			}
			result[i] = temp
			temp = 0
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("%s\n", elapsed)
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

//fmt.Print(player[i], " ")
//fmt.Println(i, " I ", j, " J ")
//if int(player[i]) < ranked[j] {
//fmt.Println(player[i], "CUK")
//temp = int32(len(ranked) + 1)
//ranked = append(ranked, int(player[i]))
//sort.Ints(ranked)
//break
//} else if int(player[i]) > ranked[j] {
//temp = int32(len(ranked) - j)
//} else if int(player[i]) == ranked[j] {
//temp = int32(j)
//break
//} else {
//temp = int32(len(ranked) - j)
//break
//}
//