package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

func WordDecoder() {
	encodedStrings := []string{"35hello7racecar5world", "0", "1612 abc", "46coding7example5civic9interview"}

	for _, v := range encodedStrings {
		processString(v)
	}
}

func processString(s string) {
	count, _ := strconv.Atoi(string(s[0]))
	fmt.Printf("There are %d word(s)\n", count)
	words := []string{}
	i := 1
	for i < len(s) {
		wordLen, _ := strconv.Atoi(string(s[i]))
		word := s[i+1 : i+1+wordLen]
		words = append(words, strings.ToUpper(word))
		i += 1 + wordLen
	}

	for _, word := range words {
		fmt.Println(word)
	}
}
