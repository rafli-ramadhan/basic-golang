package main

import "fmt"

func main() {
	// Inisialisasi slice
	s := []int{10, 20, 30, 40}
	fmt.Println("Initial slice:", s)

	// Append
	s = append(s, 50)
	fmt.Println("After append:", s)

	// Akses dan ubah elemen
	fmt.Println("Element at index 2:", s[2])
	s[2] = 99
	fmt.Println("After modifying index 2:", s)

	// Sub-slice
	sub := s[1:4]
	fmt.Println("Sub-slice [1:4]:", sub)

	// Length dan capacity
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))

	// Looping dengan range
	fmt.Println("Looping:")
	for i, v := range s {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	// Menghapus elemen di index 1 (misalnya)
	indexToRemove := 1
	s = append(s[:indexToRemove], s[indexToRemove+1:]...)
	fmt.Println("After removing index 1:", s)

	// Copy slice
	dst := make([]int, len(s))
	copy(dst, s)
	fmt.Println("Copied slice:", dst)
}
