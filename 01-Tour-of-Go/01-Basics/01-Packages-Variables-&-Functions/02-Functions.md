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

## Named Return Values

When naming return values, they act as variables and are definied at the top of the block scope. The names are used to document whatever you wish to return.

A **`return`** without arguments--referred to as a _**naked return**_--outputs the named values defined at the top of the block scope.

**NOTE**: Naked return statements should be used only in short functions and can harm readability in longer functions.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  fmt.Printf("%v is split between %v and %v\n", sum, x, y)
  return
}

func main() {
  fmt.Println(split(10))
}
```

**Output**:

```txt
10 is split between 4 and 6
4 6
```
