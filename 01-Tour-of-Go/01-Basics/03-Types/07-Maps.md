# Tour of Go: Maps

A map relates **keys** to **values**.

You can think of it like an array that stores instance objects.

The zero value of a map is `nil`.
A `nil` map has no keys, nor can keys be added.

The `make` function returns a **`map`** of the given type, initialized and ready for use.

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)

  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}

```

**Output**:

```txt
{40.68433 -74.39967}
```

</br>

**Left Off [Here](<https://tour.golang.org/moretypes/19>)**