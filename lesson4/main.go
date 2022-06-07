package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := values() // случайный набор данных
	/* Если нужно принять набор данных с консоли:
	arr := []int{}
		fmt.Scan(&arr)
		var i int
		for _, err := fmt.Scan(&i); err == nil; _, err = fmt.Scan(&i) {
			arr = append(arr, i)
		}
	*/
	fmt.Printf("До сортировки:\n %v\n", arr)
	fmt.Printf("После сортировки:\n %v", sort(arr))
}

func sort(arr []int) []int {
	for i := range arr {
		j := i + 1
		for i >= 0 && j < len(arr) && arr[j] < arr[i] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
			i--
		}
	}
	return arr
}

func values() []int {
	initial := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	var max int = 100
	var min int = -100
	for i := range initial {
		initial[i] = rand.Intn(max-min) + min
	}
	return initial
}
