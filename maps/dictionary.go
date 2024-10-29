package maps

const (
  ErrorNotFound = DictionaryError("could not find the word you were looking for")
  ErrorWordExists = DictionaryError("cannot add word because it already exists")
  ErrorWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
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

func (d Dictionary) Update(word, definition string) error {
  _, err := d.Search(word)
  switch err {
  case ErrorNotFound:
    return ErrorWordDoesNotExist
  case nil:
    d[word] = definition
  default:
    return err
  }
  return nil
}

func (d Dictionary) Delete(word string) error {
  _, err := d.Search(word)
  switch err {
  case nil:
    delete(d, word)
  default:
    return err
  }
  return nil
}

