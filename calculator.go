package calculator

import "errors"

func Add(x, y int) (int, error) {
	return x + y, nil
}

func Subtract(x, y int) (int, error) {
	return x - y, nil
}

func Multiply(x, y int) (int, error) {
	return x * y, nil
}

func Divide(x, y int) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return float64(x) / float64(y), nil
}
