package main

import (
  "learn-go-with-tests/dependencyInjection"
  "net/http"
  "log"
)

func main() {
  log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyInjection.GreeterHandler)))
}

