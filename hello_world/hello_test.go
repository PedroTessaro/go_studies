package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Pedro", "")
		want := "Hello, Pedro"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World' when an empty string is supplied", func(*testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Portuguese", func(*testing.T) {
		got := Hello("Pedro", "Portuguese")
		want := "Olá, Pedro"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Spanish", func(*testing.T) {
		got := Hello("Pedro", "Spanish")
		want := "Holla, Pedro"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in French", func(*testing.T) {
		got := Hello("Pedro", "French")
		want := "Bonjour, Pedro"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
