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
  if language == spanish {
    return fmt.Sprintf("%s%s!", spanishGreetingPrefix, name)
  } else if language == french {
    return fmt.Sprintf("%s%s!", frenchGreetingPrefix, name)
  }
  return fmt.Sprintf("%s%s!", englishGreetingPrefix, name)
}

func main() {
  fmt.Println(Hello("Ruth", ""))
}

