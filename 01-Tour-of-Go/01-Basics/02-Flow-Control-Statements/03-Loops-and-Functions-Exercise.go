package main

import (
    "fmt"
    "math"
)

// Newton's Method
func Newt(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - (z*z - x) / (2 * z)
    }
    return z
}

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

func main() {
    times := 10
    for i := 0; i < times; i++ {
        sqrt := Sqrt(float64(i))
        newt := Newt(float64(i))
        fmt.Println(i, "squared:")
        fmt.Println("  Sqrt:", sqrt)
        fmt.Println("  Newt:", newt)
        fmt.Println("  Difference:", math.Abs(sqrt-newt))
    }
}