package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// Numeric types
	var age int = 30
	var balance int64 = 1000000000
	var pi float64 = 3.14159
	var temperature float32 = 36.6

	// Text types
	var name string = "Ali"
	var initial byte = 'A'      // alias dari uint8
	var symbol rune = '♥'       // alias dari int32

	// Boolean
	var isActive bool = true

	// Aggregate types
	var scores [3]int = [3]int{80, 85, 90} // array
	var fruits []string = []string{"apple", "banana", "cherry"} // slice
	var user = map[string]string{ // map
		"name": "Ali",
		"role": "admin",
	}
	type Person struct {
		Name string
		Age  int
	}
	var p = Person{Name: "Ali", Age: 30}

	// ⛓️ Interface & error
	var any interface{} = "could be anything"
	var err error = errors.New("something went wrong")

	// Time
	var now time.Time = time.Now()

	// Print all values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Balance:", balance)
	fmt.Println("Pi:", pi)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Initial:", string(initial))
	fmt.Println("Symbol:", string(symbol))
	fmt.Println("Is Active:", isActive)
	fmt.Println("Scores:", scores)
	fmt.Println("Fruits:", fruits)
	fmt.Println("User map:", user)
	fmt.Println("Person struct:", p)
	fmt.Println("Any (interface{}):", any)
	fmt.Println("Error:", err)
	fmt.Println("Now (time.Time):", now.Format(time.RFC1123))
}
