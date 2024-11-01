package mocking

import (
  "testing"
  "bytes"
)

func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}
  mockSleeper := &MockSleeper{}
  Countdown(buffer, mockSleeper)
  got := buffer.String()
  want := `3
2
1
Go!`
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
  if mockSleeper.Calls != 3 {
    t.Errorf("not enough calls to sleeper, want 3 got %d", mockSleeper.Calls)
  }
}

