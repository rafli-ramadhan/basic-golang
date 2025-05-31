package main

import "fmt"

// Fungsi tanpa return
func greet(name string) {
    fmt.Println("Halo,", name)
}

// Fungsi dengan return
func add(a int, b int) int {
    return a + b
}

// Fungsi dengan multiple return
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("pembagi tidak boleh nol")
    }
    return a / b, nil
}

// Struct & method (pakai receiver)
type Person struct {
    Name string
    Age  int
}

func (p Person) SayHello() {
    fmt.Println("Halo, saya", p.Name)
}

func main() {
    greet("Budi")
    fmt.Println("2 + 3 =", add(2, 3))

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", result)
    }

    // Method
    p := Person{Name: "Andi", Age: 30}
    p.SayHello()
}
