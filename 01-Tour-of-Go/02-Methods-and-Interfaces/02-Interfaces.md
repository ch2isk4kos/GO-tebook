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

**Left Off [Here](<https://tour.golang.org/methods/1>)**