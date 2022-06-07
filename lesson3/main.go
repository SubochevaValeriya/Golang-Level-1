package main

import (
	"fmt"
	"math"
)

func main() {
	/*1. Доработать калькулятор из методички: больше операций и валидации данных (проверка деления на 0, возведение в степень, факториал) + форматирование строки при выводе дробного числа ( 2 знака после точки)
	 */
	fmt.Println("Введите два значения, а затем оператор из списка: +, -, *, /, !, **")
	var a, b float64
	var operation string
	fmt.Scan(&a, &b)
	fmt.Scan(&operation)

	fmt.Println(calc(a, b, operation))

	/*	2. Задание для продвинутых (необязательное). Написать приложение, которое ищет все простые числа от 0 до N включительно. Число N должно быть задано из стандартного потока ввода.
	 */

	fmt.Println("Введите значение, в пределах которого будет произведён поиск простых чисел.")
	var n int
	fmt.Scan(&n)

	fmt.Println(simple(n))
}

/*1. Доработать калькулятор из методички: больше операций и валидации данных (проверка деления на 0, возведение в степень, факториал) + форматирование строки при выводе дробного числа ( 2 знака после точки)
 */

func calc(a, b float64, operation string) string {
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
		return fmt.Sprintf("Ошибка! На ноль делитьнельзя.")
	} else if err != 0 {
		return fmt.Sprintf("Ошибка: незнакомая операция.")
	} else {
		return fmt.Sprintf("Ответ: %.2f", result)
	}
}

/*	2. Задание для продвинутых (необязательное). Написать приложение, которое ищет все простые числа от 0 до N включительно. Число N должно быть задано из стандартного потока ввода.
 */

func simple(n int) []int {
	result := []int{}
	for i := 2; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if i%j == 0 && i != j {
				break
			} else if i == j {
				result = append(result, i)
			}
		}
	}
	return result
}
