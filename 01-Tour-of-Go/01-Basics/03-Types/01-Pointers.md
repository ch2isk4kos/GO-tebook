# Tour of Go: Pointers

A **pointer** holds the memory address of a value.

</br>

The type **`*T`** is a pointer to a `T` value.
Its _zero value_ is `nil`.

> var p *int

</br>

The **`&`** operator generates a pointer to its operand.

> i := 24
> p = &i

</br>

The **`*`** operator denotes the pointer's underlying value.

> fmt.Println(*p)
>*p = 21

</br>

This is known as **_dereferencing_** or **_indirecting_**.
Unlike C, Go has no pointer arithmetic.

```go
package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i                                         // point to i
  fmt.Println("read i through the pointer", *p)   // read i through the pointer
  *p = 21                                         // set i through the pointer
  fmt.Println("new value of i: ", i)              // see the new value of i

  p = &j                                          // point to j
  *p = *p / 37                                    // divide j through the pointer
  fmt.Println("new value of j: ", j)              // see the new value of j
}
```

**Output**:

```txt
read i through the pointer:   42
new value of i:               21
new value of j:               73
```
