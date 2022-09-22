package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	fmt.Println(k/2, b/2)
	fmt.Println(bill, k, b)
	splitBill := removeIndex(bill, k)
	var totalSplitBill int32
	for _, value := range splitBill {
		totalSplitBill += value
	}
	halfBill := totalSplitBill / 2
	if halfBill == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Printf("%d\n", b-halfBill)
	}
}

func removeIndex(s []int32, index int32) []int32 {
	return append(s[:index], s[index+1:]...)
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 12)

}
