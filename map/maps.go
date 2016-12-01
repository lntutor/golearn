package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["hello"] = Vertex{40.213, -323.23}
	fmt.Println(m["hello"])
}
