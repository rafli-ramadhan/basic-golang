package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Struct yang akan dikonversi ke/dari JSON
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// Fungsi untuk menyimpan slice Product ke file JSON
func saveToJSONFile(filename string, products []Product) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // membuat output lebih rapi
	return encoder.Encode(products)
}

// Fungsi untuk membaca file JSON dan mengubahnya jadi slice Product
func readFromJSONFile(filename string) ([]Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	return products, err
}

func main() {
	// Data awal
	products := []Product{
		{ID: 1, Name: "Keyboard", Price: 250000, Stock: 12},
		{ID: 2, Name: "Mouse", Price: 150000, Stock: 8},
	}

	filename := "products.json"

	// Simpan ke file
	err := saveToJSONFile(filename, products)
	if err != nil {
		fmt.Println("Gagal menyimpan:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke", filename)

	// Baca kembali dari file
	loadedProducts, err := readFromJSONFile(filename)
	if err != nil {
		fmt.Println("Gagal membaca:", err)
		return
	}

	fmt.Println("Data dibaca dari file:")
	for _, p := range loadedProducts {
		fmt.Printf("- %s (Rp%.2f) — stok: %d\n", p.Name, p.Price, p.Stock)
	}
}
