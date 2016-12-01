package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	P := &v
	P.X = 10
	fmt.Println(v)
}
