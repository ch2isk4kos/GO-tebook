# Tour of Go: Structs

A **`struct`** is a collection of fields.
They are similar to objects in JavaScript.

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{2, 4})
}
```

**Output**:

```txt
{2 4}
```

</br>

## Struct Fields

Struct fields are accessed using dot notation.

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{2, 2}
  v.Y = 4
  fmt.Println(v.Y)
}
```

**Output**:

```txt
4
```

</br>

## Pointers to Structs

Struct fields can be accessed through a struct pointer.

To access the field of a struct to a pointer, you  could write `(*p).X`

However, the language permits us instead to write just `p.X`, without the explicit dereference.

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
```

**Output**:

```txt
{1000000000 2}
```

</br>

## Struct Literals

A **struct literal** denotes a newly allocated struct value by listing the values of its fields.

```go
package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2}           // has type Vertex
  v2 = Vertex{X: 1}           // Y:0 is implicit
  v3 = Vertex{}               // X:0 and Y:0
  p  = &Vertex{1, 2}          // has type *Vertex
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
```

**Output**:

```txt
{1 2} &{1 2} {1 0} {0 0}
```
