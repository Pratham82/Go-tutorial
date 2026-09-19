package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleQ := 0
	sauceQ := 0.0

	for _, v := range layers {
		switch v {
		case "noodles":
			noodleQ += 50
		case "sauce":
			sauceQ += 0.2
		}
	}

	return noodleQ, sauceQ

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngedient := friendsList[len(friendsList)-1]

	myList[len(myList)-1] = secretIngedient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i, v := range quantities {
		scaledQuantities[i] = v * (float64(portions) / 2)
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
