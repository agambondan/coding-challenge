package main

import "fmt"

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var result = make([]int32, len(grades))
	for i := 0; i < len(grades); i++ {
		var mod int32
		mod = grades[i] % 5
		if mod < 3 || grades[i] < 38 {
			result[i] = grades[i]
		} else if mod >= 3 {
			fmt.Println(mod, grades[i])
			if mod%2 == 0 {
				result[i] = grades[i] + 1
			} else {
				result[i] = grades[i] + 2
			}
		}
	}
	return result
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33, 72}))
}
