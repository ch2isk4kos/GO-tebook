# Tour of Go: For

The **`for`** loop is Golangs only looping construct.

A basic `for` loop is made up of 3 components seperated by semi-colons.

> the **init** statement: executed before the first iteration
> the **condition** expression: evaluated before every iteration
> the **post** statement: executed at the end of every iteration

The loop will stop iterating once the boolean condition evaluates to false.

**NOTE**: Unlike languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the **`for`** statement and the braces `{ }` are always required.

</br>

```go
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
```

**Output**:

```txt
45
```

</br>

**Left off [here](<https://tour.golang.org/flowcontrol/1>)**
