package main

import "fmt"

/*
 * Complete the 'circularArrayRotation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER k
 *  3. INTEGER_ARRAY queries
 */

func rotate(nums []int32, k int32) []int32 {
	// if rotation less than 0 or length number less equal 0 then return
	if k < 0 || len(nums) == 0 {
		return nums
	}
	r := len(nums) - int(k)%len(nums)
	nums = append(nums[r:], nums[:r]...)

	return nums
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var result []int32
	a = rotate(a, k)
	for i := 0; i < len(queries); i++ {
		result = append(result, a[queries[i]])
	}
	return result
}

func main() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}
