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

## Variables with Initializers

A  **`var`** decalaration can include one initilaier per variable. If an initializer is present, the type can be omitted and the variable will take the type of the initializer.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
  var g, js, java = "yes", true, "no!"
  fmt.Println(i, j, g, js, java)
}
```

**Output**:

```txt
1 2 yes true NO!
```

</br>

## Short Variable Declaration

You can only use the short variable declaration inside of a function in place of a **`var`** declaration with an implicit type.

```go
package main

import "fmt"

func main() {
  var i, j int = 1, 2
  k := 3

  fmt.Println(i, j, k)
}
```

**Output**:

```txt
1 2 3
```
