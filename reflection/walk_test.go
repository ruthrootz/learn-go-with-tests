package walk

import (
  "testing"
  "reflect"
)

func TestWalk(t *testing.T) {
  cases := []struct {
    Name          string
    Input         interface{}
    ExpectedCalls []string
  }{
    {
      "struct with one string field",
      struct {
        Name string
        City string
      }{"Chris", "London"},
      []string{"Chris", "London"},
    },
    {
      "struct with one string and one int field",
      struct {
        Name string
        Age int
      }{"Chris", 33},
      []string{"Chris"},
    },
    {
      "nested fields",
      struct {
        Name string
        Profile struct {
          Age int
          City string
        }
      }{"Chris", struct {
        Age int
        City string
      }{33, "London"}},
      []string{"Chris", "London"},
    },
  }

  for _, test := range cases {
    t.Run(test.Name, func(t *testing.T) {
      var got []string
      walk(test.Input, func(input string) {
        got = append(got, input)
      })

      if !reflect.DeepEqual(got, test.ExpectedCalls) {
        t.Errorf("got %v, want %v", got, test.ExpectedCalls)
      }
    })
  }
}

