package main

import "fmt"

var GlobalFlag string

func main() {
	//var chanasu = make(chan string)
	//go chanasu <- func() {
	//	fmt.Println("asu")
	//}
	var asu map[string]float64
	price := asu["MSFT"]
	fmt.Println(price)
	//now := time.Now()
	//now.Add()
	//now.Sub()
	print("[" + GlobalFlag + "]")
	go fmt.Println("1")

}
