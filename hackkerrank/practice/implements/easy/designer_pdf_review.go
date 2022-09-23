package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var charIndexNum, result []int32
	var lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(lowerCharSet); i++ {
		for j := 0; j < len(word); j++ {
			if word[j] == lowerCharSet[i] {
				charIndexNum = append(charIndexNum, int32(i))
			}
		}
	}
	for i := 0; i < len(charIndexNum); i++ {
		result = append(result, h[charIndexNum[i]])
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return int32(len(word)) * result[len(result)-1]
}

func main() {
	var h = []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	var word = "abc"
	fmt.Println(designerPdfViewer(h, word))
	var h1 = []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	var word1 = "zaba"
	fmt.Println(designerPdfViewer(h1, word1))
}
