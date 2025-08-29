package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 1. Membaca environment variable
	user := os.Getenv("USER")
	fmt.Println("USER env:", user)

	// 2. Menulis dan membaca file
	fileName := "contoh.txt"

	// Tulis file
	err := os.WriteFile(fileName, []byte("Ini contoh isi file\n"), 0644)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}

	// Baca file
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}
	fmt.Println("Isi file:", string(content))

	// 3. Membuat direktori baru
	err = os.Mkdir("folder_baru", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Gagal membuat folder:", err)
		return
	}
	fmt.Println("Folder 'folder_baru' dibuat")

	// 4. Mengecek apakah file ada
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File tidak ditemukan")
	} else {
		fmt.Println("File ditemukan:", fileName)
	}

	// 5. Rename file
	err = os.Rename(fileName, "file_rename.txt")
	if err != nil {
		fmt.Println("Gagal rename:", err)
	} else {
		fmt.Println("File berhasil di-rename ke 'file_rename.txt'")
	}

	// 6. Menghapus file
	err = os.Remove("file_rename.txt")
	if err != nil {
		fmt.Println("Gagal menghapus file:", err)
	} else {
		fmt.Println("File dihapus")
	}

	// 7. Mendapatkan working directory
	dir, _ := os.Getwd()
	fmt.Println("Working directory:", dir)

	// 8. Menyusuri direktori
	fmt.Println("Isi folder_baru:")
	filepath.Walk("folder_baru", func(path string, info os.FileInfo, err error) error {
		fmt.Println("->", path)
		return nil
	})

	// 9. Membaca argumen command-line
	args := os.Args
	fmt.Println("Argumen:", args)

	// 10. Exit dengan kode
	fmt.Println("Program selesai dengan exit code 0")
	os.Exit(0) // kode 0 = sukses
}
