package main

import "fmt"

// Contoh struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Pointer ke Integer
	a := 10
	pInt := &a
	fmt.Println("Integer value:", a)
	fmt.Println("Pointer to integer:", pInt)
	fmt.Println("Value via pointer:", *pInt)

	// Ubah nilai lewat pointer
	*pInt = 20
	fmt.Println("Updated integer value:", a)

	// 2. Pointer ke String
	str := "hello"
	pStr := &str
	fmt.Println("\nString value:", str)
	fmt.Println("Pointer to string:", pStr)
	fmt.Println("Value via pointer:", *pStr)

	// Ubah nilai lewat pointer
	*pStr = "world"
	fmt.Println("Updated string value:", str)

	// 3. Pointer ke Boolean
	flag := true
	pBool := &flag
	fmt.Println("\nBoolean value:", flag)
	fmt.Println("Pointer to boolean:", pBool)
	fmt.Println("Value via pointer:", *pBool)

	// Ubah nilai lewat pointer
	*pBool = false
	fmt.Println("Updated boolean value:", flag)

	// 4. Pointer ke Struct
	person := Person{Name: "Alice", Age: 30}
	pPerson := &person
	fmt.Println("\nStruct value:", person)
	fmt.Println("Pointer to struct:", pPerson)
	fmt.Println("Value via pointer:", *pPerson)

	// Ubah nilai lewat pointer
	pPerson.Age = 31
	fmt.Println("Updated struct value:", person)

	// 5. Pointer ke Map
	m := map[string]int{"x": 1, "y": 2}
	pMap := &m
	fmt.Println("\nMap value:", m)
	fmt.Println("Pointer to map:", pMap)
	fmt.Println("Value via pointer:", *pMap)

	// Ubah nilai lewat pointer
	(*pMap)["x"] = 100
	fmt.Println("Updated map value:", m)

	// 6. Pointer ke Slice
	slice := []string{"apple", "banana"}
	pSlice := &slice
	fmt.Println("\nSlice value:", slice)
	fmt.Println("Pointer to slice:", pSlice)
	fmt.Println("Value via pointer:", *pSlice)

	// Ubah nilai lewat pointer
	(*pSlice)[0] = "orange"
	fmt.Println("Updated slice value:", slice)

	// 7. Pointer ke Array
	arr := [3]int{1, 2, 3}
	pArr := &arr
	fmt.Println("\nArray value:", arr)
	fmt.Println("Pointer to array:", pArr)
	fmt.Println("Value via pointer:", *pArr)

	// Ubah nilai lewat pointer
	(*pArr)[1] = 20
	fmt.Println("Updated array value:", arr)
}
