package main

import (
	"context"
	"fmt"
	"time"
)

func doWorkWithTimeout(ctx context.Context) {
	fmt.Println("Work started for context with timeout")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed for context with timeout")
	case <-ctx.Done():
		fmt.Println("Work canceled for context with timeout:", ctx.Err())
	}
}

func doWorkWithDeadline(ctx context.Context) {
	fmt.Println("Work started for context with deadline")

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed for context with deadline")
	case <-ctx.Done():
		fmt.Println("Work canceled for context with deadline:", ctx.Err())
	}
}

func chooseDeadline(option string) time.Time {
	if option == "a" {
		return time.Date(2025, 6, 1, 15, 0, 0, 0, time.Local)
	} else {
		return time.Now().Add(1 * time.Hour)
	}
}

func chooseTimeout(option string) time.Duration {
	if option == "a" {
		return 3*time.Second
	} else {
		return 1*time.Hour
	}
}

func main() {
	// WithDeadline: Deadline pada waktu tertentu, misal pukul 15:00:00
	ctx1, cancel1 := context.WithDeadline(context.Background(), chooseDeadline("a"))
	defer cancel1()

	// WithTimeout: Timeout 3 detik dari waktu sekarang
	ctx2, cancel2 := context.WithTimeout(context.Background(), chooseTimeout("a"))
	defer cancel2()

	doWorkWithDeadline(ctx1)
	doWorkWithTimeout(ctx2)

	// Contoh penggunaan context dengan value
	ctxWithValue := context.WithValue(context.Background(), "userID", 12345)
	userID := ctxWithValue.Value("userID").(int)
	fmt.Println("User ID from context:", userID)
}
