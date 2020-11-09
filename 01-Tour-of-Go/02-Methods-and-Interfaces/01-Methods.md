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

## Pointer Receivers

You can declare methods with **pointer receivers**; meaning the _receiver type_ has the literal syntax `*T` for some `type T`.

( Also, `T` cannot itself be a pointer such as `*int` )

Methods with _pointer receivers_ can modify the value to which the receiver points.

Since methods often need to modify their receiver, _pointer receivers_ are more common than value receivers.

With a value receiver, the method operates on a copy of the original type value.
( This is the same behavior as for any other function argument. )

The method must have a pointer receiver to change the type value declared in the main function.

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

// notice the pointer to a type
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

// notice the value of a type
func (v Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(10)
  fmt.Println(v.Abs())
}
```

**Output**:

```txt
pointer to a type:  50
value of a type:    5
```

</br>

**Left Off [Here](<https://tour.golang.org/methods/1>)**
