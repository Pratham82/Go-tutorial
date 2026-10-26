package hamming

import "errors"

func Distance(a, b string) (int, error) {
	diff := 0

	if len(a) != len(b) {
		return 0, errors.New("Length should be equal")
	}

	for i := 0; i < len(a) && i < len(b); i++ {
		if string(a[i]) != string(b[i]) {
			diff++
		}
	}

	return diff, nil
}
