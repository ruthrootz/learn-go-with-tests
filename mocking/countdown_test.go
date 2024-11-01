package mocking

import (
  "testing"
  "bytes"
  "reflect"
)

func TestCountdown(t *testing.T) {

  t.Run("prints 3 to Go!", func(t *testing.T) {
    buffer := &bytes.Buffer{}
    Countdown(buffer, &MockSleeper{})
    got := buffer.String()
    want := `3
2
1
Go!`
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

  t.Run("sleep before every print", func(t *testing.T) {
    mockCountdownOperations := &MockCountdownOperations{}
    Countdown(mockCountdownOperations, mockCountdownOperations)
    want := []string{
      write,
      sleep,
      write,
      sleep,
      write,
      sleep,
      write,
    }
    if !reflect.DeepEqual(want, mockCountdownOperations.Calls) {
      t.Errorf("got %q want %q", mockCountdownOperations.Calls, want)
    }
  })

}

