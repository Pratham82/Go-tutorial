package raindrops

import "strconv"

var sounds = []struct {
	divisor int
	sound   string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	res := ""
	for _, s := range sounds {
		if number%s.divisor == 0 {
			res += s.sound
		}
	}
	if res == "" {
		return strconv.Itoa(number)
	}
	return res
}
