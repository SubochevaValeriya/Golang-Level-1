package forTesting

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name          string
	InputA        int
	InputB        int
	Expected      int
	ExpectedError error
}

func TestSquareOfRectangle(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 10, 10", InputA: 10, InputB: 10, Expected: 100, ExpectedError: nil},
		{Name: "case with 3, 90", InputA: 3, InputB: 90, Expected: 270, ExpectedError: nil},
		{Name: "case with -3, -1", InputA: -3, InputB: -1, Expected: 3, ExpectedError: errNegativeNumber},
		{Name: "case with -3, 0", InputA: 0, InputB: 0, Expected: 0, ExpectedError: errZero},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result, err := SquareOfRectangle(cse.InputA, cse.InputB)
			if !errors.Is(err, cse.ExpectedError) {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if err == nil {
				assert.Equalf(t, cse.Expected, result, "for %d and %d expected %d, got %d", cse.InputA, cse.InputB, cse.Expected, result)
				return
			}

			if err != nil {
				assert.Equalf(t, cse.ExpectedError, err, err.Error())
			}
		})
	}
}
