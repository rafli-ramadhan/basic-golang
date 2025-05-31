package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. Membaca dari string menggunakan io.Reader
	input := "Hello from io.Reader!"
	reader := strings.NewReader(input)

	buffer := make([]byte, 8) // buffer 8 byte
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}

	// 2. Menulis ke os.Stdout menggunakan io.Writer
	writer := os.Stdout
	_, err := writer.Write([]byte("\nMenulis ke stdout via io.Writer\n"))
	if err != nil {
		fmt.Println("Write error:", err)
	}

	// 3. Menyalin data dari reader ke writer menggunakan io.Copy
	src := strings.NewReader("Ini contoh io.Copy\n")
	dst := os.Stdout
	io.Copy(dst, src)

	// 4. Membaca semua isi reader sekaligus
	r := strings.NewReader("Baca semua data sekaligus")
	all, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("ReadAll error:", err)
	}
	fmt.Println("ReadAll:", string(all))

	// 5. io.MultiReader: gabungkan beberapa reader
	r1 := strings.NewReader("Part1 ")
	r2 := strings.NewReader("Part2 ")
	combined := io.MultiReader(r1, r2)
	io.Copy(os.Stdout, combined)
	fmt.Println()

	// 6. io.MultiWriter: tulis ke banyak writer sekaligus
	file, _ := os.Create("output.txt")
	defer file.Close()
	multiWriter := io.MultiWriter(os.Stdout, file)
	multiWriter.Write([]byte("Ditulis ke stdout dan file\n"))
}
