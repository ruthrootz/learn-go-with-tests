package racer

import (
  "time"
  "net/http"
)

func Racer(a, b string) (winner string) {
  select {
  case <- ping(a):
    return a
  case <- ping(b):
    return b
  }
}

func ping(url string) chan struct{} {
  chan := make(chan struct{})
  go func() {
    http.Get(url)
    close(chan)
  }()
  return chan
}

