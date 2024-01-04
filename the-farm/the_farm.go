package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	fodder, err := fc.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodder * fatteningFactor / float64(numCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (ce *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ce.cows, ce.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}
	return nil
}
