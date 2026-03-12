/*
Implement WordCount.
It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := map[string]int{}
	words := strings.Fields(s)

	for _, c := range words {
		_, ok := wordCount[string(c)]

		if ok {
			wordCount[string(c)] = wordCount[string(c)] + 1
		} else {
			wordCount[string(c)] = 1
		}

	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
