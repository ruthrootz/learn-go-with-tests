package maps

import (
  "errors"
)

var ErrorNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(searchTerm string) (string, error) {
  result, ok := d[searchTerm]
  if !ok {
    return "", ErrorNotFound
  }
  return result, nil
}

