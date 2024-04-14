package hello

import "testing"

func Test_hello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := hello("stas", "stas")
		want := "hello, stas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to no one", func(t *testing.T) {
		got := hello("", "")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hello("Elodie", "French")
		want := "bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
