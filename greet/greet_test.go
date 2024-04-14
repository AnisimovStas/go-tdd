package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Stas")

	got := buffer.String()
	want := "Hello, Stas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
