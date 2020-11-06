# Tour of Go: Defer

A **`defer`** statement delays the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
  defer fmt.Println("Kakos")
  fmt.Println("Chris")
}
```

**Output**:

```txt
Chris
Kakos
```

</br>

## Stacking Defers

Deferred function calls are pushed onto a stack. 
When a function returns, its deferred calls are executed in **last-in-first-out** order.

```go
package main

import "fmt"

func main() {
  fmt.Println("counting down")

  for i := 1; i <= 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("Happy New Year!")
}
```

**Output**:

```txt
counting down
Happy New Year!
10
9
8
7
6
5
4
3
2
1
```

**Andrew Gerrand: [Defer, Panic, and Recover](<https://blog.golang.org/defer-panic-and-recover>)**
