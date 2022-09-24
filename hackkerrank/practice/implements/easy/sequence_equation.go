package main

import "fmt"

/*
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func permutationEquation(p []int32) []int32 {
	// Write your code here
	var result []int32
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			if int32(i+1) == p[j] {
				for k := 0; k < len(p); k++ {
					if p[p[k]-1] == int32(i+1) {
						result = append(result, int32(k+1))
					}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(permutationEquation([]int32{2, 3, 1}))
	fmt.Println(permutationEquation([]int32{5, 2, 1, 3, 4}))
	fmt.Println(permutationEquation([]int32{2, 5, 11, 10, 1, 14, 7, 3, 16, 9, 8, 6, 18, 12, 15, 17, 13, 4}))
}
