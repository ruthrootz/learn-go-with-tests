package structsMethodsInterfaces

import (
  "testing"
)

func TestPerimeter(t *testing.T) {
  r := Rectangle {10.0, 10.0}
  got := r.Perimeter()
  want := 40.0

  if got != want {
    t.Errorf("got %.2f, want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {

  areaTests := []struct {
    shape Shape
    hasArea  float64
  }{
    {Rectangle {12, 6}, 72.0},
    {Circle {10}, 314.1592653589793},
    {Triangle {12, 6}, 36.0},
  }

  for _, tt := range areaTests {
    got := tt.shape.Area()
    if got != tt.hasArea {
      t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
    }
  }

}

