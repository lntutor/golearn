package main

import "fmt"
import "math"

func square(x float64) float64 {
  if x < 0 {
    return math.Sqrt(-x)
  }
  return math.Sqrt(x)
}

func main()  {
  fmt.Println(square(12312))
}
