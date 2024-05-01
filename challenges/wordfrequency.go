package challenges

import (
	"fmt"
	"sort"
	"strings"
)

func CountWords(text string) map[string]int {
	text = strings.ReplaceAll(strings.ReplaceAll(text, ",", ""), ".", "")

	res := make(map[string]int)

	textSplit := strings.Split(text, " ")

	for _, v := range textSplit {
		_, ok := res[v]
		if ok {
			res[v]++
		} else {
			res[v] = 1
		}
	}

	return res
}

type Result struct {
	Key   string
	Value int
}

func Top5Words(wordmap map[string]int) []Result {

	var entries []Result
	for k, v := range wordmap {
		entries = append(entries, Result{
			Key:   k,
			Value: v,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Value > entries[j].Value
	})
	return entries[0:6]
}

func WordFrequencyMain() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
