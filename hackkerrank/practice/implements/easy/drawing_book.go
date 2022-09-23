package main

import "fmt"

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
int n: the number of pages in the book
int p: the page number to turn to
*/

func pageCount(n int32, p int32) int32 {
	// Write your code here
	// jumlah halaman saat ini - jumlah halaman yang saat ini ingin di buka
	// halaman 11 pengen ke halaman 4 itu buka berapa lembar
	var result int32
	front := p / 2
	back := (n - p) / 2
	if n%2 == 0 {
		back = (n + 1 - p) / 2
	}
	if front < back {
		result = front
	} else {
		result = back
	}
	return result
}

func main() {
	fmt.Println(pageCount(5, 4))
	fmt.Println(pageCount(6, 2))
	fmt.Println(pageCount(5, 1))
}
