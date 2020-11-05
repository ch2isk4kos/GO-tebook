# Tour of Go: Functions

Go Functions can take a zero or more arguments.

```go
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(10, 24))
}
```

**Output**:

```txt
34
```
