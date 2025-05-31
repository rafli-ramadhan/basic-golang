package main

import (
	"fmt"
	"regexp"
)

// Validasi email
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// Validasi nomor telepon Indonesia
func isValidIndoPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^(?:\+62|62|0)8[1-9][0-9]{6,9}$`)
	return phoneRegex.MatchString(phone)
}

// Cek hanya huruf dan spasi
func isAlphaSpace(s string) bool {
	alphaSpace := regexp.MustCompile(`^[A-Za-z\s]+$`)
	return alphaSpace.MatchString(s)
}

// Validasi password kuat
func isStrongPassword(password string) bool {
	strongPass := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`)
	return strongPass.MatchString(password)
}

// Ambil semua angka dari string
func extractNumbers(s string) []string {
	digitRegex := regexp.MustCompile(`\d+`)
	return digitRegex.FindAllString(s, -1)
}

// Pisahkan kata (tanpa tanda baca)
func splitWords(s string) []string {
	splitWords := regexp.MustCompile(`[ ,.!?]+`)
	return splitWords.Split(s, -1)
}

// Validasi tanggal format YYYY-MM-DD
func isValidDate(date string) bool {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return dateRegex.MatchString(date)
}

// Cek awalan "Go"
func startsWithGo(s string) bool {
	startGo := regexp.MustCompile(`^Go`)
	return startGo.MatchString(s)
}

// Cek akhiran "Lang"
func endsWithLang(s string) bool {
	endLang := regexp.MustCompile(`Lang$`)
	return endLang.MatchString(s)
}

func main() {
	fmt.Println("Validasi Email:")
	fmt.Println("user@example.com:", isValidEmail("user@example.com"))  // true
	fmt.Println("user@@example:", isValidEmail("user@@example"))      // false

	fmt.Println("\nValidasi Nomor Telepon Indonesia:")
	fmt.Println("081234567890:", isValidIndoPhone("081234567890"))    // true
	fmt.Println("+6281234567890:", isValidIndoPhone("+6281234567890"))// true
	fmt.Println("123456:", isValidIndoPhone("123456"))                // false

	fmt.Println("\nCek hanya huruf dan spasi:")
	fmt.Println(`"Hello World":`, isAlphaSpace("Hello World"))        // true
	fmt.Println(`"Hello123":`, isAlphaSpace("Hello123"))              // false

	fmt.Println("\nValidasi Password Kuat:")
	fmt.Println("Aa1@abcd:", isStrongPassword("Aa1@abcd"))            // true
	fmt.Println("abcdefg:", isStrongPassword("abcdefg"))              // false

	fmt.Println("\nAmbil semua angka dari string:")
	numbers := extractNumbers("Order #12345 has 3 items, total 250000 IDR")
	fmt.Println(numbers)                                               // [12345 3 250000]

	fmt.Println("\nPisahkan kata (tanpa tanda baca):")
	words := splitWords("Hello, Gopher! How are you?")
	fmt.Println(words)                                                 // [Hello Gopher How are you]

	fmt.Println("\nValidasi Tanggal Format YYYY-MM-DD:")
	fmt.Println("2024-12-31:", isValidDate("2024-12-31"))              // true
	fmt.Println("31-12-2024:", isValidDate("31-12-2024"))              // false

	fmt.Println("\nCek Awalan dan Akhiran:")
	fmt.Println(`startsWithGo("Golang")`, startsWithGo("Golang"))      // true
	fmt.Println(`endsWithLang("Golang")`, endsWithLang("Golang"))      // true
}
