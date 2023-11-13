package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdowntStart = 3

func Countdown(w io.Writer) {
	for i := countdowntStart; i > 0; i-- {
		time.Sleep(time.Second * 1)
		fmt.Fprintln(w, i)
	}
	time.Sleep(time.Second * 1)
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
