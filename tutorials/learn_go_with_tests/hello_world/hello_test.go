package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Spanish version", func(t *testing.T) {
		got := Hello("Tony", "Spanish")
		want := "Hola, Tony"

		assertCorrectMessage(t, got, want)
	})

	t.Run("French version", func(t *testing.T) {
		got := Hello("Terry", "French")
		want := "Bonjure, Terry"

		assertCorrectMessage(t, got, want)
	})

}
