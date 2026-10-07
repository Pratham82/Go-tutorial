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

	for _, user := range lines {
		// stringMatch := re.MatchString(user)
		stringMatch := re.FindStringSubmatch(user)

		// if stringMatch {
		// 	userString := re.FindString(user)
		// 	fmt.Println(userString)
		// 	user = "[USR] " + user
		// 	test = append(test, user)
		// } else {
		// 	test = append(test, user)
		// }
		if stringMatch != nil {
			userName := stringMatch[1]
			user = fmt.Sprintf("[USR] %s ", userName) + user
			test = append(test, user)
		} else {
			test = append(test, user)
		}
	}
	return test

}
