package greets

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Pedro")

	got := buffer.String()
	want := "Hello, Pedro"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
