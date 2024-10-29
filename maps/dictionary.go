package maps

const (
  ErrorNotFound = DictionaryErr("could not find the word you were looking for")
  ErrorWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
  return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
  result, ok := d[word]
  if !ok {
    return "", ErrorNotFound
  }
  return result, nil
}

func (d Dictionary) Add(word, definition string) error {
  _, err := d.Search(word)
  switch err {
  case ErrorNotFound:
    d[word] = definition
  case nil:
    return ErrorWordExists
  default:
    return err
  }
  return nil
}

