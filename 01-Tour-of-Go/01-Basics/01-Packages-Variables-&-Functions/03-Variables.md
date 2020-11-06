# Tour of Go: Variables

The **`var`** statement declares a list of variables.
A  **`var`** statement can be declared at package or function level.

```go
package main

import "fmt"

var g, j, r string

func main() {
  g := "Go"
  j := "JavaScript"
  r := "Ruby"

  var num int   // zero value

  fmt.Println(num, g, j, r)
}
```

**Output**:

```txt
0 Go JavaScript Ruby
```

</br>

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**