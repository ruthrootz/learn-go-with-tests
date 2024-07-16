package main

import (
  "testing"
)

func TestHello(t *testing.T) {
  got := Hello("Test")
  want := "Hello, Test!"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

