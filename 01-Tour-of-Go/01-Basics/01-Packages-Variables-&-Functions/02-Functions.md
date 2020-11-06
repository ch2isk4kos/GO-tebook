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

</br>

Whenever two or more function parameters share the same type, you can omit redundency by explicitly declaring the type identifier after the last parameter:

```go
func add(x int, y int) int {
  return x + y
}

// is equivalent to...

func add(x, y int) int {
  return x + y
}
```

</br>

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**
