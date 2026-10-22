package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return n, errors.New("n must be positive")
	}

	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			steps++
		} else {
			n = n*3 + 1
			steps++
		}
	}

	return steps, nil
}
