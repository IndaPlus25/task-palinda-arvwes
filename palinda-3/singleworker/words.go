package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	text = strings.ToLower(text)
	var builder strings.Builder
	for _, c := range text {
		if !unicode.IsPunct(c) && !unicode.IsSymbol(c) {
			builder.WriteRune(c)
		}
	}
	data := builder.String()
	for _, word := range strings.Fields(data) {
		if freqs[word] == 0 {
			freqs[word] = 1
		} else {
			freqs[word] += 1
		}
	}
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	content, err := os.ReadFile("loremipsum.txt")
	if err != nil {
		return
	}
	data := string(content)

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
