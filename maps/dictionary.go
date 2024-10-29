package maps

import (

)

func Search(dictionary map[string]string, searchTerm string) string {
  return dictionary[searchTerm]
}

