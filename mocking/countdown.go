package mocking

import (
  "io"
  "fmt"
  "time"
)

var countdownStart = 3
var finalWord = "Go!"

type Sleeper interface {
  Sleep(duration time.Duration)
}

type MockSleeper struct {
  Calls int
}

func (s *MockSleeper) Sleep(duration time.Duration) {
  s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
    sleeper.Sleep(1 * time.Second)
  }
  fmt.Fprintf(out, finalWord)
}

