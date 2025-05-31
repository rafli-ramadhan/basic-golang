package main

import "fmt"

func main() {
    // Deklarasi variabel
    var name string = "Budi"
    age := 25                      // tipe otomatis (int)
    var height float64 = 175.5
    var isStudent bool = true

    // Operasi aritmatika
    age += 1       // sama seperti age = age + 1
    height -= 0.5
    age++          // increment
    age--          // decrement

    // Print variabel
    fmt.Println("Nama:", name)
    fmt.Println("Umur:", age)
    fmt.Println("Tinggi:", height)
    fmt.Println("Pelajar:", isStudent)
}
