package mocking

import (
  "io"
  "fmt"
  "time"
)

var countdownStart = 3
var finalWord = "Go!"

func Countdown(out io.Writer) {
  for i := countdownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
    time.Sleep(1 * time.Second)
  }
  fmt.Fprintf(out, finalWord)
}

