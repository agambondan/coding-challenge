package main

import "fmt"

func main() {
	intChan := make(chan int32)
	go func() {
		intChan <- 3
	}()

	fmt.Println("DI BLOCK CUY")

	n := <-intChan
	fmt.Println(n)
	// close the channel if you don't want to deadlock
	// if you comment the close function, it will be deadlock
	close(intChan)
	n = <-intChan
	fmt.Println(n)

	fmt.Println("UDAH GA DI BLOCK CUY")
}
