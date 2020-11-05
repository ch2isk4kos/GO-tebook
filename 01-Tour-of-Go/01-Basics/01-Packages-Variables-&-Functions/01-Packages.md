# Tour of Go: Packages

</br>

Each program begins running in the package main:

> **package** main

The import identifier groups together each dependency that you're abstracting:

> **import** ("fmt" "math/rand" "time)

An identifier is exported if it starts with a capital letter:

> fmt.**P**rintln( )

</br>

**Random Number Generator Example:**

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
