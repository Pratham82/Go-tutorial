package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cows int) (float64, error) {
	fodderAmt, err := f.FodderAmount(cows)
	if err != nil {
		return fodderAmt, err
	}

	fatteningFact, err := f.FatteningFactor()

	if err != nil {
		return fatteningFact, err
	}

	food := (float64(fodderAmt) * fatteningFact)
	res := food / float64(cows)
	return res, err
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	val, err := DivideFood(f, cows)

	if err != nil {
		return val, err
	}

	return val, err
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	cows          int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.customMessage)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:          cows,
			customMessage: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows:          0,
			customMessage: "no cows don't need food",
		}
	}
	return nil

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
