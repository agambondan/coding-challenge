package main

import "fmt"

/*
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func permutationEquation(p []int32) []int32 {
	// Write your code here
	var result []int32
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			if int32(i+1) == p[j] {
				fmt.Println(j+1, p[j])
				for k := 0; k < len(p); k++ {
					if p[p[k]-1] == int32(i+1) {
						result = append(result, int32(k+1))
					}
				}
			}
		}
	}
	return result
}

func main() {
	// loop until len array (i = 0; i < len(array))
	// index x = i
	fmt.Println(permutationEquation([]int32{2, 3, 1}))
	fmt.Println(permutationEquation([]int32{5, 2, 1, 3, 4}))
	//fmt.Println(permutationEquation([]int32{2, 5, 11, 10, 1, 14, 7, 3, 16, 9, 8, 6, 18, 12, 15, 17, 13, 4}))
}

// First Try
//func permutationEquation(p []int32) []int32 {
//	// Write your code here
//	fmt.Println(p, len(p))
//	var result []int32
//	for i := 0; i < len(p); i++ {
//		for j := 0; j < len(p); j++ {
//			if i+1 == int(p[j]) {
//				asem := p[p[j]-1]
//				fmt.Print("| ", i+1, j, p[j], asem, p[asem-1], " ")
//				//fmt.Println(p[p[p[j]-1]])
//				result = append(result, p[asem-1])
//			}
//		}
//	}
//	fmt.Print("|")
//	fmt.Println()
//	return result
//}

// Second Try
//func permutationEquation(p []int32) []int32 {
//	// Write your code here
//	var result []int32
//	for i := 0; i < len(p); i++ {
//		px := p[p[i]-1]
//		//fmt.Println(px)
//		for j := 0; j < len(p); j++ {
//			if px == p[j] {
//				result = append(result, px)
//			}
//		}
//	}
//	return result
//}
