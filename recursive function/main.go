package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("5! =", factorial(5)) // Output: 120

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
