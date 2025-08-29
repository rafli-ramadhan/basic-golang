package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestWriteAndReadCSV(t *testing.T) {
	filename := "test_data.csv"

	// Data yang akan ditulis
	expectedUsers := []User{
		{Name: "Andi", Email: "andi@example.com", Age: "22"},
		{Name: "Maya", Email: "maya@example.com", Age: "27"},
	}

	// Tulis data ke file CSV
	err := writeCSV(filename, expectedUsers)
	if err != nil {
		t.Fatalf("Gagal menulis ke file CSV: %v", err)
	}

	// Baca kembali file CSV
	actualUsers, err := readCSV(filename)
	if err != nil {
		t.Fatalf("Gagal membaca file CSV: %v", err)
	}

	// Bandingkan hasilnya
	if !reflect.DeepEqual(expectedUsers, actualUsers) {
		t.Errorf("Data tidak sesuai:\nGot: %#v\nWant: %#v", actualUsers, expectedUsers)
	}

	// Bersihkan file uji coba
	os.Remove(filename)
}

func TestReadCSVWithMissingFile(t *testing.T) {
	_, err := readCSV("nonexistent_file.csv")
	if err == nil {
		t.Error("Seharusnya error saat membaca file yang tidak ada")
	}
}

func TestWriteCSVCreatesFile(t *testing.T) {
	filename := "temp_output.csv"
	data := []User{{Name: "A", Email: "a@example.com", Age: "20"}}

	err := writeCSV(filename, data)
	if err != nil {
		t.Fatalf("Gagal menulis file: %v", err)
	}

	// Cek apakah file berhasil dibuat
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File %s tidak ditemukan setelah penulisan", filename)
	}

	// Cek konten minimal 1 baris data + header
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Gagal membuka file untuk pengecekan: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("Gagal membaca isi file CSV: %v", err)
	}

	if len(records) < 2 {
		t.Errorf("CSV seharusnya memiliki minimal 2 baris (header + data), got %d", len(records))
	}

	os.Remove(filename)
}
