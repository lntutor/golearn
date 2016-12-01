package main

import "fmt"

func main() {
	// q := []int{2, 3, 5, 7, 11, 13}
	// r := []bool{true, false, true, false}
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{1, false},
	}

	fmt.Println(s)
}
