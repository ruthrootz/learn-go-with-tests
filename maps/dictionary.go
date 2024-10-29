package maps

import (
  "errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(searchTerm string) (string, error) {
  result := d[searchTerm]
  if result == "" {
    return "", errors.New("could not find the word you were looking for")
  }
  return d[searchTerm], nil
}

