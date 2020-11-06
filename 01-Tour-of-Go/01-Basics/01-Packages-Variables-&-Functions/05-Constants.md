# Tour of Go: Constants

Constants are similar to variables except they are declared with the **`const`** keyword instead of `var` and cannot and intialized with _short hand declaration_ syntax.

```go
package main

import "fmt"

const Pi = 3.14

func main() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
```

**Output**:

```tzt
Hello 世界
Happy 3.14 Day
Go rules? true
```

</br>

## Numeric Constants

Numeric constants are high-precision values.

```go
package main

import "fmt"

const (
  // Create a huge number by shifting a 1 bit left 100 places.
  // In other words, the binary number that is 1 followed by 100 zeroes.
  Big = 1 << 100
  
  // Shift it right again 99 places, so we end up with 1<<1, or 2.
  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  fmt.Println("needInt(Small):", needInt(Small))
  fmt.Println("needInt(Big):", needInt(Big))
  fmt.Println("needFloat(Small):", needFloat(Small))
  fmt.Println("needInt(Big):", needFloat(Big))
}

```

**Output**:

```tzt
needInt(Small): 21
./prog.go:20:38: constant 1267650600228229401496703205376 overflows int
needFloat(Small): 0.2
needInt(Big): 1.2676506002282295e+29
```

**NOTE**: An `int` can store at maximum a **64-bit** integer, and sometimes less.

</br>

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**