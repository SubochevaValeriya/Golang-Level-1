package forTesting

import "fmt"

func ExampleMultiplication() {
	fmt.Println(Calc(3.0, 4.0, "*"))
	//Output: Ответ: 12.00
}

func ExampleDivision() {
	fmt.Println(Calc(12.0, 4.0, "/"))
	//Output: Ответ: 3.00
}

func ExampleAddition() {
	fmt.Println(Calc(3.0, 4.0, "+"))
	//Output: Ответ: 7.00
}

func ExampleSubstraction() {
	fmt.Println(Calc(3.0, 4.0, "-"))
	//Output: Ответ: -1.00
}

func ExampleDivisionByZero() {
	fmt.Println(Calc(3.0, 0.0, "/"))
	//Output: Ошибка! На ноль делить нельзя.
}
