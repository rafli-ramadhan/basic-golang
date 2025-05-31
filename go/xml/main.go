package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Struct yang akan di-encode/decode ke XML
type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:"lastName"`
	Age       int      `xml:"age"`
}

func saveToXMLFile(filename string, people []Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bungkus slice dengan root element <people>
	type PeopleWrapper struct {
		XMLName xml.Name `xml:"people"`
		People  []Person `xml:"person"`
	}

	wrapper := PeopleWrapper{People: people}

	enc := xml.NewEncoder(file)
	enc.Indent("", "  ") // supaya rapi
	return enc.Encode(wrapper)
}

func readFromXMLFile(filename string) ([]Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	type PeopleWrapper struct {
		People []Person `xml:"person"`
	}

	var wrapper PeopleWrapper
	dec := xml.NewDecoder(file)
	err = dec.Decode(&wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.People, nil
}

func main() {
	people := []Person{
		{ID: 1, FirstName: "Andi", LastName: "Wijaya", Age: 30},
		{ID: 2, FirstName: "Sari", LastName: "Nugroho", Age: 25},
	}

	filename := "people.xml"

	err := saveToXMLFile(filename, people)
	if err != nil {
		fmt.Println("Gagal menyimpan XML:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke", filename)

	loadedPeople, err := readFromXMLFile(filename)
	if err != nil {
		fmt.Println("Gagal membaca XML:", err)
		return
	}

	fmt.Println("Data dibaca dari file:")
	for _, p := range loadedPeople {
		fmt.Printf("- ID: %d, Nama: %s %s, Umur: %d\n", p.ID, p.FirstName, p.LastName, p.Age)
	}
}
