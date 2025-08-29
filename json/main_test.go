package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestSaveAndReadJSONFile(t *testing.T) {
	filename := "test_products.json"

	// Data yang ingin disimpan dan diuji
	expected := []Product{
		{ID: 1, Name: "Laptop", Price: 15000000, Stock: 5},
		{ID: 2, Name: "Monitor", Price: 2500000, Stock: 10},
	}

	// Simpan ke file
	err := saveToJSONFile(filename, expected)
	if err != nil {
		t.Fatalf("Gagal menyimpan JSON: %v", err)
	}

	// Baca kembali dari file
	result, err := readFromJSONFile(filename)
	if err != nil {
		t.Fatalf("Gagal membaca JSON: %v", err)
	}

	// Bandingkan data
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Data yang dibaca tidak sesuai:\nGot: %#v\nWant: %#v", result, expected)
	}

	// Hapus file setelah test
	_ = os.Remove(filename)
}

func TestReadFromInvalidJSONFile(t *testing.T) {
	// Buat file JSON tidak valid
	filename := "invalid.json"
	_ = os.WriteFile(filename, []byte(`invalid json`), 0644)

	_, err := readFromJSONFile(filename)
	if err == nil {
		t.Error("Seharusnya error saat membaca JSON tidak valid")
	}

	_ = os.Remove(filename)
}

func TestJSONFormat(t *testing.T) {
	filename := "format_test.json"
	products := []Product{
		{ID: 1, Name: "Phone", Price: 3000000, Stock: 7},
	}

	err := saveToJSONFile(filename, products)
	if err != nil {
		t.Fatalf("Gagal menyimpan JSON: %v", err)
	}

	// Baca isi file mentah sebagai string
	raw, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Gagal membaca file JSON langsung: %v", err)
	}

	// Cek apakah JSON bisa di-decode ulang
	var result []Product
	if err := json.Unmarshal(raw, &result); err != nil {
		t.Errorf("File tidak valid JSON: %v", err)
	}

	_ = os.Remove(filename)
}
