package main

import "fmt"

// Interface bernama Shape yang mendefinisikan method Area()
type Shape interface {
	Area() float64
}

// Struct Rectangle yang akan mengimplementasi interface Shape
type Rectangle struct {
	Width, Height float64
}

// Implementasi method Area() untuk Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Struct Circle yang juga mengimplementasi interface Shape
type Circle struct {
	Radius float64
}

// Implementasi method Area() untuk Circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Fungsi yang menerima parameter bertipe Shape (interface)
func printArea(s Shape) {
	fmt.Printf("Luas: %.2f\n", s.Area())
}

func main() {
	// Membuat objek Rectangle dan Circle
	r := Rectangle{Width: 5, Height: 4}
	c := Circle{Radius: 3}

	// Memanggil fungsi printArea dengan objek Rectangle dan Circle
	printArea(r) // Output: Luas: 20.00
	printArea(c) // Output: Luas: 28.27

	// Bisa juga assign ke variabel bertipe interface Shape
	var s Shape

	s = r
	fmt.Println("Luas Rectangle via interface:", s.Area())

	s = c
	fmt.Println("Luas Circle via interface:", s.Area())
}
