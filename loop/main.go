package main

import "fmt"

func main() {
    // For loop klasik
    for i := 0; i < 5; i++ {
        fmt.Println("For loop i:", i)
    }

    // While-style loop (pakai for)
    j := 0
    for j < 3 {
        fmt.Println("While-style loop j:", j)
        j++
    }

    // Range (foreach) loop
    fruits := []string{"apple", "banana", "cherry"}
    for idx, fruit := range fruits {
        fmt.Println("Fruit", idx, ":", fruit)
    }

    // Break & Continue
    for k := 0; k < 10; k++ {
        if k == 3 {
            continue
        }
        if k == 6 {
            break
        }
        fmt.Println("k:", k)
    }

    // Nested loop
    for x := 1; x <= 3; x++ {
        for y := 1; y <= 2; y++ {
            fmt.Printf("x=%d, y=%d\n", x, y)
        }
    }
}
