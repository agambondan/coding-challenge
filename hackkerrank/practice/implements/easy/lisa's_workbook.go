package main

import "fmt"

/*
 * Complete the 'workbook' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY arr
 */

func workbook(n int32, k int32, arr []int32) int32 {
	// Write your code here
	var result, page int32
	for i := 0; i < int(n); i++ {
		for j := 1; j <= int(arr[i]); j++ {
			if j%int(k) == 1 || k == 1 {
				page += 1
			}
			if int(page) == j {
				fmt.Printf("page : %d, arr[i] : %d\n", page, arr[i])
				result += 1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(workbook(5, 3, []int32{4, 2, 6, 1, 10}))
	fmt.Println(workbook(10, 5, []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}))
}
