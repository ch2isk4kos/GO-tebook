# Tour of Go: Arrays

The type **`[n]T`** is an array of n values of type `T`.

The expression below declares a variable `a` as an array of ten integers.

> var a [10]int

</br>

**NOTE**: An array's length is part of its type, so arrays cannot be resized but Go also provides a convenient way of working with arrays.

```go
package main

import "fmt"

func main() {
  var a [2]string
  
  a[0] = "Chris"
  a[1] = "Kakos"
  
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  tenths := [6]int{10, 20, 30, 40, 50, 60}
  fmt.Println(tenths)
}
```

**Output**:

```txt
Chris Kakos
[Chris Kakos]
[10 20 30 40 50 60]
```
