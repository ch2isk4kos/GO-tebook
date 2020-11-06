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

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**
