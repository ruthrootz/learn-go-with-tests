package main

import (
  "fmt"
)

const (
  spanish = "Spanish"
  french = "French"

  englishGreetingPrefix = "Hello, "
  spanishGreetingPrefix = "Hola, "
  frenchGreetingPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  return fmt.Sprintf("%s%s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
  switch language {
    case spanish:
      prefix = spanishGreetingPrefix
    case french:
      prefix = frenchGreetingPrefix
    default:
      prefix = englishGreetingPrefix
  }
  return
}

func main() {
  fmt.Println(Hello("Ruth", ""))
}

