package main

import (
	"os"
	"reflect"
	"testing"
)

func TestSaveAndReadXMLFile(t *testing.T) {
	filename := "test_people.xml"

	expected := []Person{
		{ID: 10, FirstName: "Budi", LastName: "Santoso", Age: 40},
		{ID: 20, FirstName: "Dewi", LastName: "Anggraini", Age: 35},
	}

	// Simpan ke file XML
	err := saveToXMLFile(filename, expected)
	if err != nil {
		t.Fatalf("Gagal menyimpan XML: %v", err)
	}

	// Baca dari file XML
	result, err := readFromXMLFile(filename)
	if err != nil {
		t.Fatalf("Gagal membaca XML: %v", err)
	}

	// Bandingkan hasil
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Data XML tidak sama:\nGot:  %#v\nWant: %#v", result, expected)
	}

	// Bersihkan file uji coba
	_ = os.Remove(filename)
}

func TestReadXMLFileNotExist(t *testing.T) {
	_, err := readFromXMLFile("file_tidak_ada.xml")
	if err == nil {
		t.Error("Seharusnya error saat file XML tidak ada")
	}
}
