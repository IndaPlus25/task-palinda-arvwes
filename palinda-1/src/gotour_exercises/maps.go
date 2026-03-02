package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	word_counts := make(map[string]int)
	for _, word := range words {
		word_counts[word] += 1
	}

	return word_counts
}

func maps() {
	wc.Test(WordCount)
}
