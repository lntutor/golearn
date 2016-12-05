package main

import "fmt"

func sum(array []int, c chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	c <- sum
}

func main() {
	c := make(chan int)
	s := []int{3, 3, 5, 5, 1, 1, 4, 4, 5, 6, 6, 3}
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
