package forTesting

import "testing"

func BenchmarkFibonacci(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Fibonacci(15)
	}
}

func BenchmarkFibonacciWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[uint]uint{
			1: 1,
			2: 1,
		}

		FibonacciWithMap(15, m)
	}
}
