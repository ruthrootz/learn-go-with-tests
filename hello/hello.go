package main

import (
  "fmt"
)

const spanish = "Spanish"
const english = "English"
const french = "French"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  switch language {
    case spanish:
      return fmt.Sprintf("%s%s!", spanishGreetingPrefix, name)
    case french:
      return fmt.Sprintf("%s%s!", frenchGreetingPrefix, name)
    default:
      return fmt.Sprintf("%s%s!", englishGreetingPrefix, name)
  }
}

func main() {
  fmt.Println(Hello("Ruth", ""))
}

