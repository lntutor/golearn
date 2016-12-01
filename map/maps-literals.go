package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m = map[string]Vertex{
// 	"1": Vertex{1.23, -2.2},
// 	"2": Vertex{-23.2, 4.0},
// }

var m = map[string]Vertex{
	"1": {1.23, -2.2},
	"2": {Lat: -23.2},
}

func main() {
	fmt.Println(m)
}
