# Tour of Go: Interfaces

An _interface type_ is defined as a set of method signatures.
A value of _interface type_ can hold any value that implements those methods.

```go
package main

import (
  "fmt"
  "math"
)

type Abser interface {
  Abs() float64
}

func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a = f  // a MyFloat implements Abser
  a = &v // a *Vertex implements Abser

  // In the following line, v is a Vertex (not *Vertex)
  // and does NOT implement Abser.
  a = v

  fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

**Output**:

```txt
./prog.go:22:4: cannot use v (type Vertex) as type Abser in assignment:
        Vertex does not implement Abser (Abs method has pointer receiver)
```

</br>

## Interfaces are Implemented Implicitly

A type implements an interface by implementing its methods. 
There is no explicit declaration of intent.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

```go
package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
  fmt.Println(t.S)
}

func main() {
  var i I = T{"Chris Kakos"}
  i.M()
}
```

**Output**:

```txt
Chris Kakos
```

</br>

## Interface Values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

> (value, type)

</br>

An **`interface`** value holds a value of a specific underlying concrete type.

Calling a method on an `interface` value executes the method of the same name on its underlying type.

```go
package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  fmt.Println(t.S)
}

type F float64

func (f F) M() {
  fmt.Println(f)
}

func main() {
  var i I

  i = &T{"Hello"}
  describe(i)
  i.M()

  i = F(math.Pi)
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
```

**Output**:

```txt
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```

</br>

## Interfaces with `nil` Underlying Values

If the concrete value inside the interface itself is `nil`, the method will be called with a _nil receiver_.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a `nil` concrete value is itself _non-nil_.

```go
package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  if t == nil {
    fmt.Println("<nil>")
    return
  }
  fmt.Println(t.S)
}

func main() {
  var i I

  var t *T
  i = t
  describe(i)
  i.M()

  i = &T{"hello"}
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
```

**Output**:

```txt
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello
```

</br>

**Left Off [Here](<https://tour.golang.org/methods/1>)**