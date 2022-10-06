package main

import "fmt"

/*
 * Complete the 'howManyGames' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER d
 *  3. INTEGER m
 *  4. INTEGER s
 */

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var result, count int32
	for i := 0; i < int(s); i++ {
		if count+p <= s {
			count += p
			if p > d && p-d > m {
				p = p - d
			} else {
				p = m
			}
			result += 1
		} else {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(howManyGames(20, 3, 6, 80))
	fmt.Println(howManyGames(20, 3, 6, 85))
	fmt.Println(howManyGames(1, 100, 1, 9777))
	fmt.Println(howManyGames(16, 2, 1, 9981))
}
