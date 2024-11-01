package mocking

import (
  "io"
  "fmt"
)

func Countdown(out io.Writer) {
  fmt.Fprintf(out, "3")
  fmt.Fprintf(out, "2")
  fmt.Fprintf(out, "1")
  fmt.Fprintf(out, "Go!")
}

