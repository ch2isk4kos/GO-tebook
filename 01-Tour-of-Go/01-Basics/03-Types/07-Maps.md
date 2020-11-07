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

## Map Literals

**Map literals** are like struct literals, but the keys are required.

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408,
  },
}

func main() {
  fmt.Println(m)
}

```

**Output**:

```txt
map[{Bell Labs: 40.68433 -74.39967} Google: {37.42202 -122.08408}]
```

</br>

If the top-level type is just a type name, you can omit it from the elements of the literal.

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{ 
  "Bell Labs": {40.68433, -74.39967},
  "Google":    {37.42202, -122.08408},
}

func main() {
  fmt.Println(m)
}
```

**Output**:

```txt
map[{Bell Labs: 40.68433 -74.39967} Google: {37.42202 -122.08408}]
```

</br>

## Mutating Maps

Insert or update an element in map  `m`:

> m [ key ] = elem

Retrieve an element:

> elem = m [ key ]

Delete an element:

> delete( m, key )

est that a key is present with a two-value assignment:

> elem, ok = m [ key ]

If key is in `m`, **`ok`** is true. 
If not, **`ok`** is false.

If key is not in the map, then `elem` is the _zero value_ for the map's element type.

**NOTE**: If `elem` or `ok` have not yet been declared you could use a short declaration form:

> elem, ok := m [ key ]

</br>

```go
package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Reeses"] = 10
  fmt.Println("The value:", m["Reeses"])

  m["Reeses"] = 24
  fmt.Println("The value:", m["Reeses"])

  fmt.Printf("m: %v\n", m)

  delete(m, "Reeses")
  fmt.Println("The value:", m["Reeses"])

  v, ok := m["Reeses"]
  fmt.Println("The value:", v, "\tPresent?", ok)
  fmt.Printf("m: %v", m)
}
```

**Output**:

```txt
The value: 0
The value: 0
m: map[Reeses:24]
The value: 0
The value: 0 	Present? false
m: map[]
```

</br>

**Left Off [Here](<https://tour.golang.org/moretypes/19>)**