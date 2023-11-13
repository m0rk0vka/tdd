package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	b := &bytes.Buffer{}

	Countdown(b)

	got := b.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
