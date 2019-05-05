package main

import "testing"

func TestPerimiter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.00, 6.00}
	got := Area(rectangle)
	want := 72.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
