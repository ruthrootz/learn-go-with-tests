package mocking

import (
  "io"
  "fmt"
  "time"
)

var countdownStart = 3
var finalWord = "Go!"

const write = "write"
const sleep = "sleep"

type Sleeper interface {
  Sleep()
}

type MockSleeper struct {
  Calls int
}

type MockCountdownOperations struct {
  Calls []string
}

func (c *MockCountdownOperations) Sleep() {
  c.Calls = append(c.Calls, sleep)
}

func (c *MockCountdownOperations) Write(p []byte) (n int, err error) {
  c.Calls = append(c.Calls, write)
  return
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

func Write(out io.Writer, toPrint string) {
  fmt.Fprintln(out, toPrint)
}

