package main

import (
  "fmt"
)

const englishGreetingPrefix = "Hello, "

func Hello(name string) string {
  return fmt.Sprintf("%s%s!", englishGreetingPrefix, name)
}

func main() {
  fmt.Println(Hello("Ruth"))
}

