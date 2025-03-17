package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello", func(t *testing.T) {
		got := Hello("infernus","English")
		want := "Hello, infernus"
		assertCorrectMessage(t, got, want)
	})

	t.Run("if empty string, default it to 'world'", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // the error points to the actual test case, and with commenting this out, the error would point to the helper function.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
