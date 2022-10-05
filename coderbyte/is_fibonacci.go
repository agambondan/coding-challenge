package main

import "fmt"

func isFibonacci(num int) string {
	var result string
	var a, b, c = 0, 1, 0
	for i := 0; i < num; i++ {
		c = b
		b = a + b
		if b == num {
			result = "yes"
			break
		}
		if b > num {
			result = "no"
			break
		}
		a = c
	}
	return result
}

func main() {
	fmt.Println(isFibonacci(60))
	for i := 0; i < 6; i++ {
		for j := 6 - 1; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
