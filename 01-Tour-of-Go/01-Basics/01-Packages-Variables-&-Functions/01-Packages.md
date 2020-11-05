# Tour of Go: Packages

Every Go program is made up of **packages**. Each program begins running in the package `main` which is the entry point of your application.

```go
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Your random number is:", rand.Intn(50))
  fmt.Println("Your custom generated number is:", rando(1, 50))
}

func rando(min, max int) int {
  return min + rand.Intn(max-min)
}
```

**Output**:

```txt
Your random number is: 10
Your custom generated number is: 24
```

</br>

## Imports