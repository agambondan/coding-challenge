package main

import "fmt"

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	var result string
	var spaceSeparatedX, spaceSeparatedY int32
	spaceSeparatedX = z - x
	spaceSeparatedY = z - y
	if spaceSeparatedX < 0 {
		spaceSeparatedX = spaceSeparatedX * -1
	}
	if spaceSeparatedY < 0 {
		spaceSeparatedY = spaceSeparatedY * -1
	}
	if spaceSeparatedX == spaceSeparatedY {
		result = fmt.Sprint("Mouse C")
	} else if spaceSeparatedX < spaceSeparatedY {
		result = fmt.Sprint("Cat A")
	} else {
		result = fmt.Sprint("Cat B")
	}
	return result
}

func main() {
	//fmt.Println(catAndMouse(1, 2, 3))
	fmt.Println(catAndMouse(22, 75, 70))
}
