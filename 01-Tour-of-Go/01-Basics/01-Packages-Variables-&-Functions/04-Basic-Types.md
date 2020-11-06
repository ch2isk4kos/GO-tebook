# Tour of Go: Basic Types

Go's basic types include:

```txt
bool

string

int  int8  int16  int32  int64

uint uint8 uint16 uint32 uint64 uintptr

byte  // alias for uint8

rune  // alias for int32
      // represents a Unicode code point

float32 float64

complex64 complex128
```

</br>

The **`int`**, **`uint`**, and **`uintptr`** types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.

When you need an integer value you should use **`int`** unless you have a specific reason to use a _sized_ or _unsigned_ integer type.

</br>

```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  yesOrNo  bool       = false
  MaxValue uint64     = 1<<64 - 1
  x        complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("yesOrNo Type: %T\tValue: %v\n", yesOrNo, yesOrNo)
  fmt.Printf("MaxValue Type: %T\tValue: %v\n", MaxValue, MaxValue)
  fmt.Printf("x Type: %T\tValue: %v\n", x, x)
}
```

**Output**:

```txt
yesOrNo Type: bool    Value: false
MaxValue Type: uint64 Value: 18446744073709551615
x Type: complex128    Value: (2+3i)
```

</br>

## Zero Values

A **zero value** is the output of a variable declared without explicitly initializing its value.

> numeric types: 0
> strings: ""
> boolean: false

</br>

```go
package main

import "fmt"

func main() {
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

**Output**:

```txt
0 0 false ""
```

</br>

## Type Conversions

The expression `T(v)` converts the value `v` to the type `T`.

Unlike in C, assignment between items of a different type requires an explicit conversion.

> var i int = 42
> var f float64 = float64( i )
> var u uint = uint( f )

</br>

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
```

</br>

## Type Inference

When declaring a variable without explicity specifying its type, the variables type is determined by the initialized value on the right hand side.

If the type is explicitly defined, it infers the variable type by the value it was initailized to.

```go
package main

import "fmt"

var i int

func main() {  
  j := i
  fmt.Printf("j = %v is of type %T\n", j, j)
  j = j + 10.24
  fmt.Printf("j = %v is of type %T\n", j, j)
  
  v := 10
  fmt.Printf("v = %v is of type %T\n", v, v)
  
  w := 10.24
  fmt.Printf("v = %v is of type %T\n", w, w)
  
  y := 0.867 + 0.5i
  fmt.Printf("w = %v is of type %T\n", y, y)
}
```

**Output**:

```txt
j = 0 is of type int
./prog.go:11:8: constant 10.9i truncated to integer

v = 24 is of type int
w = 33.34 is of type float64
y = (0.867+0.5i) is of type complex128
```

</br>

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**
