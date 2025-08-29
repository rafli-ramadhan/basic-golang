package main

import "fmt"

func main() {
	// Inisialisasi map
	m := map[string]int{
		"apple":  5,
		"banana": 7,
	}
	fmt.Println("Initial map:", m)

	// Menambah/ubah elemen
	m["cherry"] = 3
	m["banana"] = 10
	fmt.Println("After add/update:", m)

	// Akses nilai
	val := m["apple"]
	fmt.Println("Value for 'apple':", val)

	// Mengecek keberadaan key
	if v, ok := m["mango"]; ok {
		fmt.Println("Mango exists with value:", v)
	} else {
		fmt.Println("Mango does not exist")
	}

	// Menghapus elemen
	delete(m, "banana")
	fmt.Println("After deleting 'banana':", m)

	// Looping
	fmt.Println("Looping map:")
	for key, value := range m {
		fmt.Printf("%s -> %d\n", key, value)
	}

	// Panjang map
	fmt.Println("Length of map:", len(m))
}
