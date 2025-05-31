package main

import "fmt"

func main() {
	// assignment in integer
	x := 10

	x += 5 // x = 15
	x -= 2 // x = 13
	x *= 2 // x = 26
	x /= 2 // x = 13
	x %= 4 // x = 1

	fmt.Println("Final value of x:", x)

	// assignment in slice
	s := []int{1, 2, 3}
	s[0] += 5      // OK, s[0] = 6
	fmt.Println(s) // [6 2 3]

	s2 := []int{4, 5}
	s = s2 // assignment slice, OK
	fmt.Println(s) // [4 5]

	// assignment in array
	var a1 = [2]int{1, 2}
	var a2 = [2]int{3, 4}
	a1 = a2 // OK, harus sama panjang dan tipe

	a1[1] *= 2      // OK, elemen bisa pakai compound
	fmt.Println(a1) // [3 8]

	// assignment in map
	m := map[string]int{"a": 1, "b": 2}
	m["a"] += 10 // OK, jadi 11

	m2 := m // assign map ke variabel lain (by reference)
	fmt.Println(m2)
}
