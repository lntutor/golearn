package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	add := func(x, y float64) float64 {
		return x + y
	}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(add(5, 4))
	fmt.Println(compute(add))
	fmt.Println(compute(hypot))
}
