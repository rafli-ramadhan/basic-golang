package main

import "fmt"

func main() {
	// 1. Deklarasi array kosong (nilai default = 0)
	var arr1 [3]int
	fmt.Println("1. Default array:", arr1)

	// 2. Inisialisasi dengan nilai
	arr2 := [3]string{"go", "lang", "array"}
	fmt.Println("2. Initialized array:", arr2)

	// 3. Inisialisasi otomatis panjangnya
	arr3 := [...]float64{3.14, 2.71, 1.41}
	fmt.Println("3. Auto-length array:", arr3)

	// 4. Akses dan ubah elemen
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("4. After update:", arr1)

	// 5. Looping dengan for
	fmt.Print("5. Loop with for: ")
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()

	// 6. Looping dengan range
	fmt.Println("6. Loop with range:")
	for idx, val := range arr2 {
		fmt.Printf("Index %d: %s\n", idx, val)
	}

	// 7. Panjang array
	fmt.Println("7. Length of arr3:", len(arr3))

	// 8. Copy array (by value)
	arrCopy := arr1
	arrCopy[0] = 999
	fmt.Println("8. Original array:", arr1)
	fmt.Println("   Copied array :", arrCopy)

	// 9. Array multidimensi
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("9. Multidimensional array:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
