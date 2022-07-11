package forTesting

import (
	"fmt"
	"math"
)

func Calc(a, b float64, operation string) string {
	var result float64
	var err, errZero int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0.0 {
			errZero = -1
		} else {
			result = a / b
		}
	case "!":
		result = 1
		if a == 0 {
			result = 1
		} else {
			var i float64 = 1
			for ; i <= a; i++ {
				result = i * result
			}
		}
	case "**":
		result = math.Pow(a, b)

	default:
		err = -1
	}
	if errZero != 0 {
		return fmt.Sprintf("Ошибка! На ноль делить нельзя.")
	} else if err != 0 {
		return fmt.Sprintf("Ошибка: незнакомая операция.")
	} else {
		return fmt.Sprintf("Ответ: %.2f", result)
	}
}
