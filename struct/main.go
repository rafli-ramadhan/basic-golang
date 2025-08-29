package main

import (
	"fmt"
	"strings"
)

// Definisi struct Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
}

// Method untuk Person: FullName mengembalikan nama lengkap
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Method untuk Person: AddHobby menambahkan hobi baru
func (p *Person) AddHobby(hobby string) {
	p.Hobbies = append(p.Hobbies, hobby)
}

func main() {
	// Membuat instance struct
	p1 := Person{
		FirstName: "Ahmad",
		LastName:  "Sulaiman",
		Age:       30,
		Hobbies:   []string{"Membaca", "Bermain Sepak Bola"},
	}

	// Mengakses field
	fmt.Println("Nama Depan:", p1.FirstName)
	fmt.Println("Umur:", p1.Age)

	// Mengakses method
	fmt.Println("Nama Lengkap:", p1.FullName())

	// Menambah hobi baru
	p1.AddHobby("Coding")
	fmt.Println("Hobi sekarang:", strings.Join(p1.Hobbies, ", "))

	// Modifikasi langsung field
	p1.Age = 31
	fmt.Println("Umur setelah update:", p1.Age)

	// Contoh penggunaan slice dan loop
	fmt.Println("Daftar hobi:")
	for i, h := range p1.Hobbies {
		fmt.Printf("%d. %s\n", i+1, h)
	}

	// Contoh penggunaan map (key-value)
	personsByAge := map[int][]Person{
		p1.Age: {p1},
	}

	// Menambahkan person baru
	p2 := Person{FirstName: "Siti", LastName: "Nur", Age: 25}
	personsByAge[p2.Age] = append(personsByAge[p2.Age], p2)

	// Menampilkan orang berdasarkan umur
	fmt.Println("\nOrang berdasarkan umur:")
	for age, persons := range personsByAge {
		fmt.Printf("Umur %d:\n", age)
		for _, person := range persons {
			fmt.Println("-", person.FullName())
		}
	}
}
