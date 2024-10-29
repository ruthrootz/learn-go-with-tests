package maps

import (
  "errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(searchTerm string) (string, error) {
  result, ok := d[searchTerm]
  if !ok {
    return "", errors.New("could not find the word you were looking for")
  }
  return result, nil
}

