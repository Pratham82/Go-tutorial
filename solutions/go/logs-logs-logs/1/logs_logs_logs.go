package logs

import "strings"

type Logs map[rune]string

var logs = Logs{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	//words := strings.Fields(log)
	foundRune := ""
	for _, word := range log {
		foundRune = logs[word]
		if foundRune != "" {
			return foundRune
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var foundRune strings.Builder
	for _, s := range log {
		if s == oldRune {
			foundRune.WriteString(string(newRune))
		} else {
			foundRune.WriteString(string(s))
		}
	}
	return foundRune.String()
}

/*
 *
 *
 if s == oldRune {
			foundRune += string(newRune)
		} else {
			foundRune += string(s)
		}
*/

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0

	for range log {
		count++
	}

	return count <= limit
}
