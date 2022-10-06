package main

import "fmt"

/*
 * Complete the 'taumBday' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER b
 *  2. INTEGER w
 *  3. INTEGER bc
 *  4. INTEGER wc
 *  5. INTEGER z
 */

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	b64, w64, bc64, wc64, z64 := int64(b), int64(w), int64(bc), int64(wc), int64(z)
	var result int64
	if bc == wc {
		result = b64*bc64 + w64*wc64
	} else if bc > wc+z {
		result = w64*wc64 + b64*wc64 + b64*z64
	} else if wc > bc+z {
		result = b64*bc64 + w64*bc64 + w64*z64
	} else {
		result = b64*bc64 + w64*wc64
	}
	return result
}

func main() {
	t := [][]int32{{10, 10, 1, 1, 1}, {5, 9, 2, 3, 4}, {3, 6, 9, 1, 1}, {7, 7, 4, 2, 1}, {3, 3, 1, 9, 2}}
	for i := 0; i < len(t); i++ {
		fmt.Println(taumBday(t[i][0], t[i][1], t[i][2], t[i][3], t[i][4]))
	}
	//18192035842
	//20582630747
	t1 := [][]int32{{27984, 1402, 619246, 615589, 247954}, {95677, 39394, 86983, 311224, 588538}}
	for i := 0; i < len(t1); i++ {
		fmt.Println(taumBday(t1[i][0], t1[i][1], t1[i][2], t1[i][3], t1[i][4]))
	}

}
