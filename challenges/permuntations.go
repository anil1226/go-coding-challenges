package challenges

import (
	"fmt"
)

func CheckPermutations(str1, str2 string) bool {
	// implement
	if len(str1) != len(str2) {
		return false
	}

	//simpler
	// for _, v := range str1 {
	// 	if !strings.ContainsRune(str2, v) {
	// 		return false
	// 	}
	// }

	str1Map := getMapofString(str1)
	str2Map := getMapofString(str2)

	if len(str1Map) != len(str2Map) {
		return false
	}

	for k, _ := range str2Map {
		if str1Map[k] != str2Map[k] {
			return false
		}
	}
	return true
}

func getMapofString(str string) map[string]int {
	str1Map := make(map[string]int)
	for _, v := range str {
		if str1Map[string(v)] > 0 {
			str1Map[string(v)] += 1
		} else {
			str1Map[string(v)] = 1
		}
	}
	return str1Map
}

func CheckPermutationsMain() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcmz"
	str2 := "medac"

	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)

}
