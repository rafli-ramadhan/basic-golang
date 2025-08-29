package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Sqrt(16):", math.Sqrt(16))
	fmt.Println("2^8:", math.Pow(2, 8))
	fmt.Println("Sin(π/2):", math.Sin(math.Pi/2))
	fmt.Println("Log(100):", math.Log10(100))
	fmt.Println("Max(3.5, 2.7):", math.Max(3.5, 2.7))
	fmt.Println("Round(3.6):", math.Round(3.6))
	fmt.Println("Mod(10.5, 3):", math.Mod(10.5, 3))
	fmt.Println("Hypot(3,4):", math.Hypot(3, 4)) // 5 (segitiga siku-siku)

	// konversi rad ke derajat
	rad := math.Pi / 2
	deg := rad * (180 / math.Pi)
	fmt.Println(rad)
	fmt.Println(deg)

	// konversi derajat ke radian
	deg2 := 90.0
	rad2 := deg2 * (math.Pi / 180)
	fmt.Println(deg2)
	fmt.Println(rad2)
}
