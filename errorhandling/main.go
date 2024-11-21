package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if x < 0 {
		return 0, errors.New("negative numerator")
	}
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func main() {
	result, err := divide(-10, 0)
	if err != nil {
		fmt.Println("Error", fmt.Errorf("division operation failed: %w", err))
	} else {
		fmt.Println("Result:", result)
	}
}
