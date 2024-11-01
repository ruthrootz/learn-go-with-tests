package mocking

import (
  "io"
  "fmt"
  "time"
)

var countdownStart = 3
var finalWord = "Go!"

type Sleeper interface {
  Sleep()
}

type MockSleeper struct {
  Calls int
}

type DefaultSleeper struct {}

func (s *MockSleeper) Sleep() {
  s.Calls++
}

func (s *DefaultSleeper) Sleep() {
  time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
    sleeper.Sleep()
  }
  fmt.Fprintf(out, finalWord)
}

