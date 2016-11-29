package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	result := 1.0
	for i:= 0; i < 100; i++ {
		result = result - (result*result - x)/(2*result)
	}
	return result;
}

func main() {
	fmt.Println(Sqrt(2))
}
