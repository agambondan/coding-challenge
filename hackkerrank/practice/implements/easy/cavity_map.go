package main

import (
	"fmt"
)

/*
 * Complete the 'cavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func replaceAtIndex1(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

func cavityMap(grid []string) []string {
	// Write your code here
	var result []string
	for i := 0; i < len(grid); i++ {
		result = append(result, grid[i])
	}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			if grid[i][j] > grid[i-1][j] && grid[i][j] > grid[i][j-1] && grid[i][j] > grid[i][j+1] && grid[i][j] > grid[i+1][j] {
				result[i] = replaceAtIndex1(result[i], 'X', j)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(cavityMap([]string{"1112", "1912", "1892", "1234"}))
	var expectedOutput = "877X35X37998356X345X3X532988235X123463X468352X4X5X3671X42X84367X2X1X3X51737X3X2865X331611X548X15"
	output := cavityMap([]string{"884179362992919143428887745218617594484248579735335683155622684818584863837361995843138587668725",
		"877935637998356834563953298823581234638468352948583671742984367929193951737839286593316118548915",
		"958485147934716746152527692875528343974751298294723783494844152444653553259494475842956576368312"})
	fmt.Println(output[1] == expectedOutput, output, expectedOutput)
	if output[1] == expectedOutput {
		fmt.Println(output[1] == expectedOutput, output[1], expectedOutput)
	}
}

//func cavityMap(grid []string) []string {
//	// Write your code here
//	var result []string
//	for i := 0; i < len(grid); i++ {
//		result = append(result, grid[i])
//	}
//	for i := 1; i < len(grid)-1; i++ {
//		for j := 1; j < len(grid)-1; j++ {
//			if (0 < i && i < len(grid)-1) && (0 < j && j < len(grid)-1) {
//				top, _ := strconv.Atoi(grid[i-1][:j])
//				bottom, _ := strconv.Atoi(grid[i+1][:j])
//				left, _ := strconv.Atoi(grid[i][j-1 : j])
//				middle, _ := strconv.Atoi(grid[i][j : j+1])
//				right, _ := strconv.Atoi(grid[i][j+1 : j+2])
//				fmt.Println(top, bottom, left, middle, right)
//				if middle > left && middle > right && middle > top && middle > bottom {
//					result[i] = replaceAtIndex1(result[i], 'X', j)
//				}
//			}
//		}
//	}
//	return result
//}
