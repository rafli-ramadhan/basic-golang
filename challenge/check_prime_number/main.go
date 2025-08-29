package main

import (
	"fmt"
	"math"
)

// cek apakah n adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false // 0,1, dan angka negatif bukan prima
	}
	// Cek pembagi dari 2 sampai sqrt(n)
	sqrtN := int(math.Sqrt(float64(n)))
	fmt.Println(sqrtN)
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	testNumbers := []int{1, 2, 3, 4, 5, 16, 17, 18, 19, 20}

	for _, num := range testNumbers {
		if isPrime(num) {
			fmt.Printf("%d is prime\n", num)
		} else {
			fmt.Printf("%d is not prime\n", num)
		}
	}
}
