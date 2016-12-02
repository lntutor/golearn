// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
//
// type Stringer interface {
//     String() string
// }

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Loi Nguyen", 27}
	fmt.Println(a)
}
