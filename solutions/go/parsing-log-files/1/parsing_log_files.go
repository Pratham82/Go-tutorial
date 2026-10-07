package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	// 	re, err := regexp.Compile(`[TRC|DBG|INF|WRN|ERR|FTL]`)
	//
	// 	if err != nil {
	// 		return false
	// 	}
	//
	// 	return re.MatchString(text)

	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)

}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	strArr := re.Split(text, -1)
	return strArr

}

func CountQuotedPasswords(lines []string) int {
	// re := regexp.MustCompile(`(?i)"*password"`)
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	count := 0

	for _, c := range lines {
		if re.MatchString(c) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile("end-of-line[0-9]*")

	newString := re.ReplaceAllString(text, "")

	return newString

}

func TagWithUserName(lines []string) []string {
	// re := regexp.MustCompile(`User\s+[a-zA-Z0-9]`)
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)

	test := []string{}

	for _, c := range lines {
		// sm := re.MatchString(c)
		sm := re.FindStringSubmatch(c)

		// if sm {
		// 	userString := re.FindString(c)
		// 	fmt.Println(userString)
		// 	c = "[USR] " + c
		// 	test = append(test, c)
		// } else {
		// 	test = append(test, c)
		// }
		if sm != nil {
			userName := sm[1]
			c = fmt.Sprintf("[USR] %s ", userName) + c

			test = append(test, c)
		} else {
			test = append(test, c)
		}
	}
	return test

}
