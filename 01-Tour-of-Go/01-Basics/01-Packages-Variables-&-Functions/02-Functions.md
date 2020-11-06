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

Whenever two or more function parameters share the same type, you can reduce redundency by explicitly declaring the type identifier after the final parameter:

```go
func add(x int, y int) int {
  return x + y
}

// is equivalent to...

func add(x, y int) int {
  return x + y
}
```

**NOTE**: the type declaration--to the right of the function parameters--prior to the block scope--defines the functions return type.

</br>

## Return Multiple Values

A function can return any number of results.

```go
package main

import "fmt"

func swapOrder(x, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swapOrder("Kakos", "Chris")
  fmt.Println(a, b)
}
```

**Output**:

```txt
Chris Kakos
```

</br>

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**
