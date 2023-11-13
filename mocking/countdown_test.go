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

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
