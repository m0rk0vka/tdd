package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	b := &bytes.Buffer{}
	s := &SpySleeper{}

	Countdown(b, s)

	got := b.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if s.Calls != 4 {
		t.Errorf("Not enough calls for sleeper, want 4 got %d", s.Calls)
	}
}
