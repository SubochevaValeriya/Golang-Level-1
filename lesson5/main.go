package main

import "fmt"

func main() {
	var a uint

	fmt.Scan(&a)

	if a < 1 {
		panic("Введите число от 1")
	}

	m := map[uint]uint{
		1: 1,
		2: 1,
	}

	fmt.Println(fibonacci(a))
	fmt.Println(fibonacciWithMap(a, m))
}

func fibonacci(a uint) uint {
	if a == 1 || a == 2 {
		return 1
	} else {
		return fibonacci(a-1) + fibonacci(a-2)
	}
}

func fibonacciWithMap(a uint, m map[uint]uint) uint {

	if a == 1 || a == 2 {
		return 1
	} else {
		if _, ok := m[a-1]; ok == false {
			m[a-1] = fibonacciWithMap(a-1, m)
		}

		if _, ok := m[a-2]; ok == false {
			m[a-2] = fibonacciWithMap(a-1, m)
		}

		return m[a-1] + m[a-2]
	}
}
