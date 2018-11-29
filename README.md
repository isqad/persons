# Persons, fake persons

## Usage

1) Copy `names_ru`, `families_ru` and `snames_ru` to your project

2)
```go
package main

import (
  "github.com/isqad/persons"
  "log"
)

func main() {
  log.Print(persons.GetPerson())
}
```
