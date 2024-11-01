package mocking

import (
  "io"
  "fmt"
)

var countdownStart = 3
var finalWord = "Go!"

func Countdown(out io.Writer) {
  for i := countdownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
  }
  fmt.Fprintf(out, finalWord)
}

