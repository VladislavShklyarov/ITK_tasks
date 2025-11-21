package main

import (
	"fmt"
	"sort"
	"strings"
)

func WordFrequency(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		result[word]++
	}
	return result
}

func PrintWordFrequency(freqMap map[string]int) {
	words := make([]string, 0, len(freqMap))

	for word := range freqMap {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		return freqMap[words[i]] > freqMap[words[j]]
	})

	for _, word := range words {
		fmt.Printf("%s: %d\n", word, freqMap[word])
	}
}

func main() {

	text := "golang is great and golang is fast"

	PrintWordFrequency(WordFrequency(text))

}
