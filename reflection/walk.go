package walk

import ()

func walk(x interface{}, fn func(string)) {
  fn("test string")
}

