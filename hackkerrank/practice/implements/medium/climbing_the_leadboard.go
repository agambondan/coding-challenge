package main

import (
	"fmt"
	"sort"
	"time"
)

// https://stackoverflow.com/questions/73838996/climbing-the-leaderboard-hackerrank-solutions-in-golang
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	ranks := ranked[:1]
	last := ranks[0]
	for _, score := range ranked[1:] {
		if score != last {
			ranks = append(ranks, score)
		}
		last = score
	}

	climb := make([]int32, 0, len(player))
	for _, score := range player {
		rank := sort.Search(
			len(ranks),
			func(i int) bool { return ranks[i] <= score },
		)
		climb = append(climb, int32(rank+1))
	}
	return climb
}

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
}
