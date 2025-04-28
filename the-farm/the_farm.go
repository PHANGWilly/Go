package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cowsNumber int) (float64, error) {
	initialFodderTotal, error := fodderCalculator.FodderAmount(cowsNumber)
	if error != nil {
		return 0, error
	}

	fatteningFactor, error := fodderCalculator.FatteningFactor()
	if error != nil {
		return 0, error
	}

	return initialFodderTotal / float64(cowsNumber) * fatteningFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cowsNumber int) (float64, error) {
	if cowsNumber <= 0 {
		return 0, errors.New("invalid number of cows")
	} else {
		return DivideFood(fodderCalculator, cowsNumber)
	}
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	number  int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.number, e.message)
}

func ValidateNumberOfCows(cowsNumber int) error {
	switch {
	case cowsNumber < 0:
		return &InvalidCowsError{
			number:  cowsNumber,
			message: "there are no negative cows",
		}
	case cowsNumber == 0:
		return &InvalidCowsError{
			number:  cowsNumber,
			message: "no cows don't need food",
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
