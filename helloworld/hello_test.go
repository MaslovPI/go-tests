package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Pavel", "")
		want := "Hello, Pavel"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		asserCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		asserCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Oondine", "French")
		want := "Bonjour, Oondine"
		asserCorrectMessage(t, got, want)
	})
}

func asserCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
