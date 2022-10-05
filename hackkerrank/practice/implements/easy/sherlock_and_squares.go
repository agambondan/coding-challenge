package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'squares' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER a
 *  2. INTEGER b
 */

func squares(a int32, b int32) int32 {
	// Write your code here
	return int32(math.Floor(math.Sqrt(float64(b))) - math.Ceil(math.Sqrt(float64(a))) + 1)
}

func main() {
	fmt.Println(squares(3, 9))
	fmt.Println(squares(3, 28))
	fmt.Println(squares(465868129, 988379794))

}
