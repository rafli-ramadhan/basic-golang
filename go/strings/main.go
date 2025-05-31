package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "  Hello, Gophers! Welcome to Go strings package.  "

	// 1. Trim spaces di depan dan belakang
	trimmed := strings.TrimSpace(text)
	fmt.Println("TrimSpace:", trimmed)

	// 2. Convert ke lowercase
	lower := strings.ToLower(trimmed)
	fmt.Println("ToLower:", lower)

	// 3. Convert ke uppercase
	upper := strings.ToUpper(trimmed)
	fmt.Println("ToUpper:", upper)

	// 4. Cek apakah string mengandung substring
	contains := strings.Contains(trimmed, "Gophers")
	fmt.Println("Contains 'Gophers':", contains)

	// 5. Mengganti substring (replace)
	replaced := strings.ReplaceAll(trimmed, "Go", "Golang")
	fmt.Println("ReplaceAll 'Go' with 'Golang':", replaced)

	// 6. Membagi string berdasarkan delimiter
	parts := strings.Split(trimmed, " ")
	fmt.Println("Split by space:", parts)

	// 7. Menggabungkan array string menjadi satu string
	joined := strings.Join(parts, "-")
	fmt.Println("Join with '-':", joined)

	// 8. Mengecek prefix dan suffix
	hasPrefix := strings.HasPrefix(trimmed, "Hello")
	hasSuffix := strings.HasSuffix(trimmed, "package.")
	fmt.Println("HasPrefix 'Hello':", hasPrefix)
	fmt.Println("HasSuffix 'package.':", hasSuffix)

	// 9. Menghitung jumlah kemunculan substring
	count := strings.Count(trimmed, "o")
	fmt.Println("Count 'o':", count)

	// 10. Mendapatkan index substring pertama
	index := strings.Index(trimmed, "Gophers")
	fmt.Println("Index of 'Gophers':", index)

	// 11. Untuk membandingkan dua string case-insensitive
	fmt.Println(strings.EqualFold("GoLang", "golang"))

	// 12. Mengulangi string sebanyak n kali
	repeated := strings.Repeat("Go! ", 3)
	fmt.Println("Repeat 'Go! ' 3 times:", repeated)

	// 13. Menghapus prefix dan suffix dari string
	trimPrefix := strings.TrimPrefix(trimmed, "Hello, ")
	trimSuffix := strings.TrimSuffix(trimmed, " package.")
	fmt.Println("TrimPrefix 'Hello, ':", trimPrefix)
	fmt.Println("TrimSuffix ' package.':", trimSuffix)

	// 14. Memecah string menjadi field (kata) berdasarkan whitespace
	fields := strings.Fields(trimmed)
	fmt.Println("Fields (split by whitespace):", fields)

	// 15. Mengganti hanya n kemunculan pertama substring
	replaceN := strings.Replace(trimmed, "o", "0", 2)
	fmt.Println("Replace first 2 'o' with '0':", replaceN)

	// 16. Mendapatkan substring berdasarkan index (gunakan slicing biasa)
	substr := trimmed[7:14] // dari indeks 7 sampai 13 (14 exclusive)
	fmt.Println("Substring trimmed[7:14]:", substr)

	// 17. Mengecek apakah string hanya berisi huruf atau angka (pakai rune dan unicode package)
	// contoh sederhana cek apakah semua karakter adalah huruf
	onlyLetters := true
	for _, r := range trimmed {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) && !strings.ContainsRune(",.!?", r) {
			onlyLetters = false
			break
		}
	}
	fmt.Println("Only letters, spaces, or punctuation:", onlyLetters)

	// 18. Membuat string builder untuk gabung string secara efisien
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("Builder!")
	fmt.Println("Using strings.Builder:", builder.String())

	// 19. SplitN: membagi string tapi hanya sebanyak N bagian
	splitN := strings.SplitN(trimmed, " ", 4)
	fmt.Println("SplitN 4 parts:", splitN)

	// 20. IndexFunc: mencari posisi rune pertama yang memenuhi kondisi
	indexFunc := strings.IndexFunc(trimmed, func(r rune) bool {
		return r == 'W' || r == 'w'
	})
	fmt.Println("IndexFunc (first 'W' or 'w'):", indexFunc)
}
