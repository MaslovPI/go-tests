package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Pavel")

	got := buffer.String()
	want := "Hello, Pavel"

	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
