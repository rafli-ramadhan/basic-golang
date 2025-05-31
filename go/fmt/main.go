package main

import (
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// 1. Print biasa (tanpa format)
	fmt.Println("Hello, Gopher!") // Dengan newline
	fmt.Print("Ini tanpa newline ")
	fmt.Print("lanjutannya\n")

	// 2. Print dengan format
	name := "Ari"
	age := 28
	fmt.Printf("Halo %s, umur kamu %d tahun\n", name, age)

	// 3. Format ke string, bukan cetak langsung
	msg := fmt.Sprintf("Halo %s, umur kamu %d tahun", name, age)
	fmt.Println("Dari Sprintf:", msg)

	// 4. Menampilkan nilai struct%
	user := User{"Budi", 30, "budi@example.com"}

	fmt.Printf("User: %v\n", user)      // value
	fmt.Printf("User: %+v\n", user)     // field names
	fmt.Printf("User: %#v\n", user)     // Go-syntax representation
	fmt.Printf("Type: %T\n", user)      // tipe data
	fmt.Printf("Alamat pointer: %p\n", &user) // pointer

	// 5. Format angka
	n := 12345.6789
	fmt.Printf("Angka: %.2f\n", n)       // float 2 desimal
	fmt.Printf("Angka: %.3f\n", n)       // float 3 desimal
	fmt.Printf("Hex: %x\n", 255)         // hexadecimal
	fmt.Printf("Binary: %b\n", 10)       // biner
	fmt.Printf("Scientific: %e\n", n)    // notasi ilmiah

	// 6. Boolean
	ok := true
	fmt.Printf("Boolean: %t\n", ok)

	// 7. Rune dan karakter
	ch := 'G'
	fmt.Printf("Karakter: %c, Unicode: %U\n", ch, ch)

	// 8. Mengambil input dari user (basic)
	var input string
	fmt.Print("Ketik sesuatu: ")
	fmt.Scanln(&input)
	fmt.Println("Kamu mengetik:", input)
}
