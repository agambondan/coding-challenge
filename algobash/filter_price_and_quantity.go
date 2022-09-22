package main

import "fmt"

// SolutionFilter is your solution code.
func SolutionFilter(N int, itemList [][]int, M int, query [][]int) {
	// Your code starts here.
	var count = make([]int, M)
	for i := 0; i < M; i++ {
		for k := 0; k < N; k++ {
			if (query[i][0] <= itemList[k][0] && itemList[k][0] <= query[i][1]) && (query[i][2] <= itemList[k][1] && itemList[k][1] <= query[i][3]) {
				count[i] += 1
			}
		}
	}
	for _, v := range count {
		fmt.Println(v)
	}
}

func main() {
	var itemList = [][]int{{5, 2}, {4, 2}, {7, 7}, {8, 3}, {4, 5}}
	var query = [][]int{{5, 5, 2, 2}, {4, 7, 5, 7}, {4, 8, 2, 7}}
	SolutionFilter(len(itemList), itemList, len(query), query)
}
