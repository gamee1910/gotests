package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Game", "English")
		want := "Hello, Game"

		assertCorrectMessages(t, got, want)
	})

	t.Run("say 'Hello' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessages(t, got, want)
	})

	t.Run("say 'Hello' in Viet Nam", func(t *testing.T) {
		got := Hello("Game", vietnam)
		want := "Xin chào, Game"

		assertCorrectMessages(t, got, want)
	})

	t.Run("say 'Hello' in French", func(t *testing.T) {
		got := Hello("Game", french)

		want := "Bonjour, Game"

		assertCorrectMessages(t, got, want)
	})

	t.Run("say 'Hello' in Spanish", func(t *testing.T) {
		got := Hello("Game", spanish)

		want := "Hola, Game"

		assertCorrectMessages(t, got, want)
	})
}

func assertCorrectMessages(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: [%q] want: [%q]", got, want)
	}
}
