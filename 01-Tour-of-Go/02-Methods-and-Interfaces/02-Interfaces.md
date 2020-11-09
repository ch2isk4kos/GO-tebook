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

## `nil` Interface Values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

```go
package main

import "fmt"

type I interface {
  M()
}

func main() {
  var i I
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
```

**Output**:

```txt
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x49a8cf]

goroutine 1 [running]:
main.main()
  /tmp/sandbox318293311/prog.go:12 +0x8f
```

</br>

## The Empty Interface

The interface type that specifies zero methods is known as the empty interface:

> interface{ }

An empty interface may hold values of any type.
( Every type implements at least zero methods. )

Empty interfaces are used by code that handles values of unknown type.
For example: `fmt.Print()` takes any number of arguments of `type interface{}`.

```go
package main

import "fmt"

func main() {
  var i interface{}
  describe(i)

  i = 42
  describe(i)

  i = "hello"
  describe(i)
}

func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}
```

**Output**:

```txt
(<nil>, <nil>)
(42, int)
(hello, string)
```

</br>

## Type Assertions

A **type assertion** provides access to an interface value's underlying concrete value.

```txt
t := i.(T)
```

This statement asserts that the interface value i holds the concrete type T and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values:

1. the underlying value
2. a boolean value that reports whether the assertion succeeded.

```txt
t, ok := i.(T)
```

If `i` holds a `T`, then `t` will be the underlying value and `ok` will be `true`.

If not, `ok` will be `false` and `t` will be the _zero value_ of `type T`, and no panic occurs.

**NOTE**: the similarity between this syntax and that of reading from a map.

```go
package main

import "fmt"

func main() {
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s)

  s, ok := i.(string)
  fmt.Println(s, ok)

  f, ok := i.(float64)
  fmt.Println(f, ok)

  f = i.(float64)     // panic
  fmt.Println(f)
}
```

**Output**:

```txt
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
  /tmp/sandbox300166282/prog.go:17 +0x1fe
```

</br>

## Type `switch`

A **type `switch`** is a construct that permits several _type assertions_ in series.

A _type switch_ is like a regular `switch` statement, but the cases in a _type switch_ specify types (not values), and those values are compared against the type of the value held by the given interface value.

```txt
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

The declaration in a _type switch_ has the same syntax as a type assertion `i.(T)`, but the specific _type T_ is replaced with the keyword `type`.

This `switch` statement tests whether the interface value `i` holds a value of type `T` or `S`.

In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`.

In the default case (where there is no match), the variable `v` is of the same _interface type_ and value as `i`.

```go
package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21)
  do("hello")
  do(true)
}
```

**Output**:

```txt
Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!
```

</br>

**Left Off [Here](<https://tour.golang.org/methods/1>)**