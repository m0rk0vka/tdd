// dependency injection
package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}
	Greet(&b, "Vova")

	got := b.String()
	want := "Hello, Vova"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
