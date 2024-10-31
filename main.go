package main

import (
  "learn-go-with-tests/dependencyInjection"
  "os"
)

func main() {
  dependencyInjection.Greet(os.Stdout, "Test")
}

