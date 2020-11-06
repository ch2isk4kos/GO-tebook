# Tour of Go: Slices

An array has a fixed size.

A **`slice`** is a dynamically-sized, flexible view into the elements of an array and much more commonly used in practice.

</br>

The type **`[]T`** is a slice with elements of type `T`.
A slice is formed by specifying two indices, a **low** and **high** bound, separated by a colon:

> a [ low : high ]

</br>
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:

> a [ 1:4 ]

</br>

```go
package main

import "fmt"

func main() {
  primes := [6]int{2, 3, 5, 7, 11, 13}

  var s []int = primes[1:4]
  fmt.Println(s)
}
```

**Output**:

```txt
[3 5 7]
```

</br>

## Slices Are Built On Top of Arrays

A **slice** does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.

```go
package main

import "fmt"

func main() {
  names := [4]string{
    "Pike",
    "Griesemer",
    "Thompson",
    "Cox",
  }
  fmt.Println(names)

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  b[0] = "---"
  fmt.Println(a, b)
  fmt.Println(names)
}
```

**Output**:

```txt
[Pike Griesemer Thompson Cox]
[Pike Griesemer] [Griesemer Thompson]
[Pike ---] [--- Thompson]
[Pike --- Thompson Cox]
```

</br>

## Slice Literals

A **slice literal** is like an array literal without the length.

This is an array literal:

> [ 3 ]bool{ true, true, false }

It creates an array, then builds a slice that references it:

> [ ]bool{ true, true, false }

</br>

```go
package main

import "fmt"

func main() {
  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  s := []struct {
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)
}
```

**Output**:

```txt
[2 3 5 7 11 13]
[true false true true false true]
[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
```

</br>

## Slice Defaults

When slicing, you may omit the **high** or **low** bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

> var a [ 10 ]int

</br>

these slice expressions are equivalent:

> a [ 0:10 ]
> a [ :10 ]
> a [ 0: ]
> a [ : ]

</br>

```go
package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}

  s = s[1:4]
  fmt.Println(s)

  s = s[:2]
  fmt.Println(s)

  s = s[1:]
  fmt.Println(s)
}
```

**Output**:

```txt
[3 5 7]
[3 5]
[5]
```

</br>

## Slice Length and Capacity

A slice has both a **length** and a **capacity**.

</br>

The **length** of a slice is the number of elements it contains.
> len( s )

The **capacity** of a slice is the number of elements in the underlying array, counting from the first element in the slice.

> cap( s  )

</br>

You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

```go
package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // Slice the slice to give it zero length.
  s = s[:0]
  printSlice(s)

  // Extend its length.
  s = s[:4]
  printSlice(s)

  // Drop its first two values.
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len = %d\tcap = %d %v\n", len(s), cap(s), s)
}
```

**Output**:

```txt
len = 6   cap = 6 [2 3 5 7 11 13]
len = 0   cap = 6 []
len = 4   cap = 6 [2 3 5 7]
len = 2   cap = 4 [5 7]
```

</br>

## `nil` Slices

The _zero value_ of a slice is `nil`.

A **nil slice** has a length and capacity of 0 and has no underlying array.

```go
package main

import "fmt"

func main() {
  var s []int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("nil!")
  }
}
```

**Output**:

```txt
[] 0 0
nil!
```

</br>

## Creating A Slice With `make`

Slices can be created with the built-in **`make`** function; this is how you create dynamically-sized arrays.

</br>

The **`make`** function allocates a zeroed array and returns a slice that refers to that array:

> a := make ( [ ]int, 5 )

_len(a) = 5_

</br>

To specify a capacity, pass a third argument to **`make`**:

> b := make ( [ ]int, 0, 5 )

_len(b) = 0, cap(b) = 5_

> b = b [ :cap( b ) ]

_len(b) = 5, cap(b) = 5_

> b = b [ 1: ]

_len(b) = 4, cap(b) = 4_

</br>

```go
package main

import "fmt"

func main() {
  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s -> len = %d\tcap = %d %v\n",
    s, len(x), cap(x), x)
}
```

**Output**:

```txt
a -> len = 5   cap = 5 [0 0 0 0 0]
b -> len = 0   cap = 5 []
c -> len = 2   cap = 5 [0 0]
d -> len = 3   cap = 3 [0 0 0]
```

</br>

## Slices of Slices ( Matrices )

Slices can contain any type, including other slices.

```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Create a tic-tac-toe board.
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  // The players take turns.
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}
```

**Output**:

```txt
X _ X
O _ X
_ _ O
```

</br>

## Appending to a Slice

It is common to append new elements to a slice, and so Go provides a built-in **`append`** function.

> func append( s [ ]T, vs ...T ) [ ]T

</br>

The first parameter `s` of **`append`** is a slice of type `T`, and the rest are `T` values to **`append`** to the slice.

The resulting value of **`append`** is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

```go
package main

import "fmt"

func main() {
  var s []int
  printSlice(s)

  // append works on nil slices.
  s = append(s, 0)
  printSlice(s)

  // The slice grows as needed.
  s = append(s, 1)
  printSlice(s)

  // We can add more than one element at a time.
  s = append(s, 2, 3, 4)
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len = %d\tcap = %d %v\n", len(s), cap(s), s)
}
```

**Output**:

```txt
len = 0   cap = 0 []
len = 1   cap = 1 [0]
len = 2   cap = 2 [0 1]
len = 5   cap = 6 [0 1 2 3 4]
```

**Andrew Gerrand: [Go Slices - Usage and Internals](<https://blog.golang.org/slices-intro>)**
