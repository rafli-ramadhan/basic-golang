package main

import "testing"

func TestFullName(t *testing.T) {
	p := Person{FirstName: "Ahmad", LastName: "Sulaiman"}
	expected := "Ahmad Sulaiman"
	if p.FullName() != expected {
		t.Errorf("FullName() = %s; want %s", p.FullName(), expected)
	}
}

func TestAddHobby(t *testing.T) {
	p := Person{FirstName: "Ahmad", LastName: "Sulaiman", Hobbies: []string{"Membaca"}}
	p.AddHobby("Coding")

	if len(p.Hobbies) != 2 {
		t.Errorf("AddHobby() failed, length = %d; want 2", len(p.Hobbies))
	}

	if p.Hobbies[1] != "Coding" {
		t.Errorf("AddHobby() failed, last hobby = %s; want Coding", p.Hobbies[1])
	}
}
