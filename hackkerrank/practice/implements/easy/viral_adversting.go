package main

import "fmt"

/*
 * Complete the 'viralAdvertising' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func viralAdvertising(n int32) int32 {
	// Write your code here
	var shared, liked, result int32 = 5, 2, 2
	for i := 2; i <= int(n); i++ {
		shared = liked * 3
		liked = shared / 2
		result += liked
	}
	return result
}

func main() {
	fmt.Println(viralAdvertising(3))
	fmt.Println(viralAdvertising(4))
	fmt.Println(viralAdvertising(50))
	fmt.Println(viralAdvertising(34))
}
