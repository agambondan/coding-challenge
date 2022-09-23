package main

import "fmt"

/*
 * Complete the 'angryProfessor' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY a
 */

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var result string
	var countStudents int32
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			countStudents += 1
		}
	}
	if countStudents >= k {
		result = "NO"
	} else {
		result = "YES"
	}
	return result
}

func main() {
	var k int32 = 3
	var a = []int32{-1, -3, 4, 2}
	fmt.Println(angryProfessor(k, a))
	var k1 int32 = 2
	var a1 = []int32{0, -1, 2, 1}
	fmt.Println(angryProfessor(k1, a1))
}
