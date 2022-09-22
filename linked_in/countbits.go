package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countBits(num uint32) int32 {
	fmt.Println(strconv.FormatInt(100, 2))
	numBitString := fmt.Sprintf("%08b", num)
	fmt.Println(numBitString, strings.Trim(numBitString, "0"), strings.Count(numBitString, "1"))
	return int32(strings.Count(numBitString, "1"))
}

func main() {
	var value uint32 = 100
	countBits(value)
}
