package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	var typeBird = make(map[string]int32)
	// Write your code here
	for i := 0; i < len(arr); i++ {
		key := strconv.Itoa(int(arr[i]))
		_, exist := typeBird[key]
		if exist {
			typeBird[key] += 1
		} else {
			typeBird[key] = 1
		}
	}
	var keys = make([]string, 0, len(typeBird))
	var values = make([]int, 0, len(typeBird))
	for key, value := range typeBird {
		keys = append(keys, key)
		values = append(values, int(value))
	}
	sort.Strings(keys)
	sort.Ints(values)
	var result, temp int32
	for i, key := range keys {
		atoi, _ := strconv.Atoi(key)
		if i == 0 {
			fmt.Println(key, typeBird[key], "ASO", i)
			result = int32(atoi)
			temp = typeBird[key]
		} else if typeBird[key] > temp {
			if result < int32(atoi) {
				result = int32(atoi)
				temp = typeBird[key]
				fmt.Println(key, "HMMM")
			}
		}
	}
	fmt.Println(typeBird)
	return result
}

func main() {
	var arr = []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4, 5, 5, 5, 5}
	fmt.Println(migratoryBirds(arr))

}
