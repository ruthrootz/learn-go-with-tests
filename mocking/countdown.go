package mocking

import (
  "io"
  "fmt"
)

func Countdown(out io.Writer) {
  fmt.Fprintf(out, "3\n2\n1\nGo!")
}

