package chance

import (
	"math/rand/v2"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	var min, max = 1, 20

	res := min + rand.IntN(max-min+1)

	return res

}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {

	return rand.Float64() * 12.0

}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}
