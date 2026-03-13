package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {

	words := strings.Fields(text)
	var wg sync.WaitGroup
	numWorkers := 8
	freqCh := make(chan map[string]int, numWorkers)
	var chunks [][]string
	chunkSize := len(words) / numWorkers

	for i := 0; i < len(words); i += chunkSize {
		j := i + chunkSize
		if j > len(words) {
			j = len(words)
		}
		chunks = append(chunks, words[i:j])
	}

	for _, chunk := range chunks {
		wg.Add(1)
		go func(textPart []string) {
			defer wg.Done()
			localFreqs := make(map[string]int)

			// clean and count words localy
			for _, word := range textPart {
				var builder strings.Builder

				//clean the words chars
				for _, c := range word {
					if !unicode.IsPunct(c) && !unicode.IsSymbol(c) {
						builder.WriteRune(unicode.ToLower(c))
					}
				}
				word = builder.String()
				localFreqs[word]++
			}
			freqCh <- localFreqs
		}(chunk)
	}
	//wait for wait group, then close
	go func() {
		wg.Wait()
		close(freqCh)
	}()
	// count total of word counts
	totalFreqs := make(map[string]int)
	for freqs := range freqCh {
		for word, count := range freqs {
			totalFreqs[word] += count
		}
	}
	return totalFreqs
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
