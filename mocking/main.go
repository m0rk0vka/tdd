package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdowntStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdowntStart; i > 0; i-- {
		s.Sleep()
	}
	for i := countdowntStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	ds := &DefaultSleeper{}
	Countdown(os.Stdout, ds)
}
