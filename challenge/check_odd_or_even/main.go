package main

import "fmt"

func oddOrEven(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	num := 10
	fmt.Printf("%d is %s\n", num, oddOrEven(num))

	num = 7
	fmt.Printf("%d is %s\n", num, oddOrEven(num))
}
