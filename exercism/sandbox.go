package main

import "fmt"

type Logs map[rune]string

var logs = Logs{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	foundRune := ""
	fmt.Println(foundRune)
	for _, s := range log {
		if s == oldRune {
			foundRune += string(newRune)
		} else {
			foundRune += string(s)
		}

	}
	return foundRune
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0

	for range log {
		count++
	}

	return count <= limit
}

func main() {
	fmt.Println(WithinLimit("hello❗", 6))

}
