package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'happyLadybugs' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING b as parameter.
 */

// happyLadybugs
func happyLadybugs(b string) string {
	// Write your code here
	var temp []string
	for i := 0; i < len(b); i++ {
		s := string(b[i])
		if s != "_" {
			count := strings.Count(b, s)
			if count < 2 {
				return "NO"
			} else {
				for j := 0; j < count; j++ {
					temp = append(temp, s)
				}
			}
		}
	}
	if strings.Count(b, "_") > 0 {
		return "YES"
	} else {
		var pair int
		for i := 0; i < len(b)-1; i++ {
			if b[i] == b[i+1] {
				pair += 1
			} else if pair > 0 {
				pair = 0
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	T := []string{"AABBC", "AABBC_C", "_", "DD__FQ_QQF", "AABCBC"}
	for i := 0; i < len(T); i++ {
		fmt.Println(happyLadybugs(T[i]))
	}
	fmt.Println()
	T = []string{"RBY_YBR", "X_Y__X", "__", "B_RRBR"}
	for i := 0; i < len(T); i++ {
		fmt.Println(happyLadybugs(T[i]))
	}
	fmt.Println()
	T = []string{"_", "RBRB", "RRRR", "BBB", "BBB_"}
	for i := 0; i < len(T); i++ {
		fmt.Println(happyLadybugs(T[i]))
	}
	fmt.Println()
	T = []string{"G", "GR", "_GR_", "_R_G_", "R_R_R", "RRGGBBXX", "RRGGBBXY"}
	//T = []string{"RRGGBBXX"}
	for i := 0; i < len(T); i++ {
		fmt.Println(happyLadybugs(T[i]))
	}
	fmt.Println()
	T = []string{"_FWYSSENEDBO_KSEVUAB_WZ_GASASVEVS_O_NSVBYFNADE_WWVSBKAE_F_VAS_F_AAAEWBE_WEAEOAYV", "ZBF_MIFUXJNQGQRFZVRQUFFFFNGFIBJ_XZVIRFGMJRJFVMNJMF",
		"YFCA_NXMGJYYGCMMGGGXYNAMFNJJX_", "CBLJUKUWSTIIUKUBQSITSULTJKCUSKBCKB", "A_TOJRPRW__JOJP__WAJT",
	}
	for i := 0; i < len(T); i++ {
		fmt.Println(happyLadybugs(T[i]))
	}
}
