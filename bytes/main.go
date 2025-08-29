package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("  Hello, Gophers! Welcome to Go.  ")

	// 1. Trim whitespace
	trimmed := bytes.TrimSpace(data)
	fmt.Println("TrimSpace:", string(trimmed))

	// 2. Convert ke lowercase
	lower := bytes.ToLower(trimmed)
	fmt.Println("ToLower:", string(lower))

	// 3. Cek contains
	fmt.Println("Contains 'Go':", bytes.Contains(trimmed, []byte("Go")))

	// 4. Replace
	replaced := bytes.ReplaceAll(trimmed, []byte("Go"), []byte("Golang"))
	fmt.Println("Replace 'Go' with 'Golang':", string(replaced))

	// 5. Split dan Join
	parts := bytes.Split(trimmed, []byte(" "))
	fmt.Printf("Split: %q\n", parts)
	joined := bytes.Join(parts, []byte("-"))
	fmt.Println("Join with '-':", string(joined))

	// 6. Prefix dan Suffix
	fmt.Println("HasPrefix 'Hello':", bytes.HasPrefix(trimmed, []byte("Hello")))
	fmt.Println("HasSuffix '.':", bytes.HasSuffix(trimmed, []byte(".")))

	// 7. Equal dan EqualFold
	fmt.Println("EqualFold 'GoLang' & 'golang':", bytes.EqualFold([]byte("GoLang"), []byte("golang")))

	// 8. Menggunakan bytes.Buffer
	var b bytes.Buffer
	b.WriteString("Hello")
	b.Write([]byte(", World!"))
	fmt.Println("Buffer content:", b.String())
}
