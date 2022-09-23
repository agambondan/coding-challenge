package main

import "fmt"

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	var output string
	if x1 == x2 && v1 == v2 {
		output = "YES"
	} else if x1 == x2 && v1 != v2 {
		output = "NO"
	} else {
		if v1 > v2 && ((x2-x1)%(v1-v2)) == 0 {
			output = "YES"
		} else {
			output = "NO"
		}
	}
	return output
}

func main() {
	fmt.Println(kangaroo(0, 2, 5, 3))
}
