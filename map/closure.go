package main

import "fmt"

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func trippleAdder(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x + y + z
		}
	}
}

func main() {
	add5, add10 := adder(5), adder(10)
	fmt.Println(add5(3))
	fmt.Println(add10(7))
	fmt.Println(adder(5)(3))
	fmt.Println(trippleAdder(1)(2)(3))

}
