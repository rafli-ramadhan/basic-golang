package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Struct untuk menampung data user
type User struct {
	Name  string
	Email string
	Age   string
}

// Fungsi untuk membaca file CSV
func readCSV(filename string) ([]User, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []User
	for _, row := range rows[1:] { // Lewati header
		if len(row) < 3 {
			continue
		}
		users = append(users, User{
			Name:  row[0],
			Email: row[1],
			Age:   row[2],
		})
	}

	return users, nil
}

// Fungsi untuk menulis ke file CSV
func writeCSV(filename string, users []User) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Tulis header
	writer.Write([]string{"Name", "Email", "Age"})

	// Tulis data
	for _, user := range users {
		err := writer.Write([]string{user.Name, user.Email, user.Age})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	filename := "data.csv"

	// Baca file CSV
	users, err := readCSV(filename)
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	fmt.Println("Data dari file:")
	for i, user := range users {
		fmt.Printf("%d. %s - %s - %s\n", i+1, user.Name, user.Email, user.Age)
	}

	// Tambahkan data baru
	newUser := User{Name: "Rina", Email: "rina@example.com", Age: "28"}
	users = append(users, newUser)

	// Simpan kembali ke file
	err = writeCSV(filename, users)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}

	fmt.Println("\nData berhasil diperbarui dan disimpan ke file.")
}
