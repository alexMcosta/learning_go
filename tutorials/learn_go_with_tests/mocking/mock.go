package main

import (
	"fmt"
	"io"
	"os"
)

const startNum = 3
const finalWord = "Go!"

func Countdown(out io.Writer) {
	for i := startNum; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
