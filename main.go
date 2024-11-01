package main

import (
  "learn-go-with-tests/dependencyInjection"
  "learn-go-with-tests/mocking"
  "net/http"
  "log"
  "os"
)

func main() {
  writeToStdOut()
}

func writeToStdOut() {
  sleeper := &mocking.DefaultSleeper{}
  mocking.Countdown(os.Stdout, sleeper)
}

func runServer() {
  log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyInjection.GreeterHandler)))
}

