package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdowntStart = 3

func Countdown(w io.Writer) {
	for i := countdowntStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
