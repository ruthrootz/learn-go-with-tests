package structsMethodsInterfaces

type Rectangle struct {
  Width float64
  Height float64
}

type Circle struct {
  Radius float64
}

func (r Rectangle) Perimeter() float64 {
  return r.Width * 2 + r.Height * 2
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (c Circle) Area() float64 {
  return c.Radius * c.Radius * 3.141592653589793
}

