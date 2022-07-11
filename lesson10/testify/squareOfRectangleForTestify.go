package forTesting

import "errors"

var errNegativeNumber = errors.New("length can't be less than 0")
var errZero = errors.New("length of rectangle can't be equal to 0")

func SquareOfRectangle(a, b int) (int, error) {
	var err error
	if a < 0 || b < 0 {
		err = errNegativeNumber
	}

	if a == 0 || b == 0 {
		err = errZero
	}

	result := a * b

	return result, err
}
