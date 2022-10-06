package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */

func acmTeam(topic []string) []int32 {
	// Write your code here
	var temp []int32
	var result = make([]int32, 2)
	for i := 0; i < len(topic); i++ {
		for j := i + 1; j < len(topic); j++ {
			var tempValue int32
			for k := 0; k < len(topic[i]); k++ {
				if string(topic[i][k]) == "1" || string(topic[j][k]) == "1" {
					tempValue += 1
				}
			}
			temp = append(temp, tempValue)
			tempValue = 0
		}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	for i := 0; i < len(temp); i++ {
		if temp[i] == temp[len(temp)-1] {
			result[1] += 1
		}
	}
	result[0] = temp[len(temp)-1]
	return result
}

func main() {
	topic := []string{"10101", "11100", "11010", "00101"}
	fmt.Println(acmTeam(topic))
}
