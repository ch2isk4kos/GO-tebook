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

## Methods and Pointer Indirection

functions with a pointer argument must take a pointer:

```txt
var v Vertex
ScaleFunc(v, 5)     // Compile error!
ScaleFunc(&v, 5)    // OK
```

While methods with pointer receivers take either a value or a pointer as the receiver when they are called:

```txt
var v Vertex
v.Scale(5)          // OK
p := &v
p.Scale(10)         // OK
```

```go
package main

import "fmt"

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(2)
  ScaleFunc(&v, 10)

  p := &Vertex{4, 3}
  p.Scale(3)
  ScaleFunc(p, 8)

  fmt.Println(v, p)
}
```

**Output**:

```txt
{60 80} &{96 72}
```

**NOTE**: For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the Scale method has a pointer receiver.

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```txt
var v Vertex
fmt.Println(AbsFunc(v))         // OK
fmt.Println(AbsFunc(&v))        // Compile error!
```

While methods with value receivers take either a value or a pointer as the receiver when they are called:

```txt
var v Vertex
fmt.Println(v.Abs())            // OK
p := &v
fmt.Println(p.Abs())            // OK
```

</br>

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

func AbsFunc(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
  fmt.Println(AbsFunc(v))

  p := &Vertex{4, 3}
  fmt.Println(p.Abs())
  fmt.Println(AbsFunc(*p))
}
```

**Output**:

```txt
5
5
5
5
```

**NOTE**: the method call `p.Abs()` is interpreted as `(*p).Abs()`.

</br>

## Choosing a Value or Point Receiver

There are two reasons to use a pointer receiver:

1. so the method can modify the value that its receiver points to.
2. to avoid copying the value on each method call.

All methods on a given type should have either value or pointer receivers, but not a mixture of both.

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := &Vertex{3, 4}
  fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
  v.Scale(5)
  fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

**Output**:

```txt
Before scaling: &{X:3 Y:4}, Abs: 5
After scaling: &{X:15 Y:20}, Abs: 25
```
