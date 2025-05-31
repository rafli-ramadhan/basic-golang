package main

import "testing"

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 5, Height: 4}
	want := 20.0
	got := r.Area()
	if got != want {
		t.Errorf("Rectangle.Area() = %v; want %v", got, want)
	}
}

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 3}
	want := 3.14159 * 3 * 3
	got := c.Area()
	if got != want {
		t.Errorf("Circle.Area() = %v; want %v", got, want)
	}
}

func TestPrintArea(t *testing.T) {
	// Karena printArea hanya print ke stdout, kita tidak bisa tangkap outputnya
	// secara langsung tanpa import package lain. Jadi test ini hanya memastikan
	// printArea tidak panic saat dipanggil.
	r := Rectangle{Width: 2, Height: 3}
	c := Circle{Radius: 1}

	printArea(r)
	printArea(c)
}
