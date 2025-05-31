package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Waktu sekarang
	now := time.Now()
	fmt.Println("Sekarang:", now)

	// 2. Format waktu
	formatted := time.Now().Format("02-01-2006 15:04:05")
	fmt.Println("Format custom:", formatted)

	// 3. Parsing string ke time.Time
	parsedTime, err := time.Parse("2006-01-02", "2025-06-01")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("Parsed time:", parsedTime)
	}

	// 4. Tambah dan kurang durasi
	nextWeek := time.Now().AddDate(0, 0, 7)
	fmt.Println("7 hari lagi:", nextWeek)

	pastHour := time.Now().Add(-1 * time.Hour)
	fmt.Println("1 jam lalu:", pastHour)

	// 5. Durasi antara dua waktu
	duration := time.Now().Sub(pastHour)
	fmt.Println("Durasi sejak 1 jam lalu:", duration)

	duration2 := time.Since(pastHour)
	fmt.Println("Durasi sejak 1 jam lalu:", duration2)

	// 6. Sleep selama 2 detik
	fmt.Println("Tunggu 2 detik...")
	time.Sleep(2 * time.Second)
	fmt.Println("Lanjut!")

	// 7. Timezone
	local := time.Now().Local()
	utc := time.Now().UTC()
	fmt.Println("Local time:", local)
	fmt.Println("UTC:", utc)

	// 8. Timer (untuk dijalankan di goroutine)
	fmt.Println("Timer 1 detik...")
	timer := time.NewTimer(1 * time.Second)
	<-timer.C
	fmt.Println("Timer selesai")

	// 9. Menggunakan time.After
	fmt.Println("Menggunakan time.After...")
	<-time.After(1 * time.Second)
	fmt.Println("Selesai setelah time.After")

	// 10. Perbandingan waktu
	t1 := time.Date(2025, 6, 1, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("t1 Before t2:", t1.Before(t2))
	fmt.Println("t1 After t2:", t1.After(t2))
	fmt.Println("t1 Equal t2:", t1.Equal(t2))

	// 11. Ambil komponen waktu
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	fmt.Printf("Tanggal: %d-%s-%d\n", day, month, year)
	fmt.Printf("Waktu: %02d:%02d:%02d\n", hour, min, sec)
}
