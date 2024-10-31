package dependencyInjection

import (
  "testing"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "Ruth")
  got := buffer.String()
  want := "Hello, Ruth"
  if got != want {
    t.Errorf("got %q want %q", got want)
  }
}

