# Tour of Go: Methods

Go does not have classes. You can however define methods on types.

A **method** is a function with a special receiver argument.

The receiver appears in its own argument list between the `func` keyword and the _method name_.

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
}
```

**Output**:

```txt
5
```

**NOTE**: the Abs method has a receiver of `type Vertex` named `v`.

</br>

## Methods are Functions

A method is just a function with a receiver argument.
You can declare a method on non-struct types, too.

You can only declare a method with a receiver whose type is defined in the same package as the method.

You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

```go
package main

import (
  "fmt"
  "math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())
}
```

**Output**:

```txt
1.4142135623730951
```

</br>

**Left Off [Here](<https://tour.golang.org/methods/1>)**
