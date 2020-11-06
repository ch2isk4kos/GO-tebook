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

**Rob Pike: [Go Declaration Syntax](https://blog.golang.org/declaration-syntax)**