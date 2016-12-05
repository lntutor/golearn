package main

import (
	"fmt"
	"time"
)

func sendMsgToChannel1(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second * 2)
	}
}

func sendMsgToChannel2(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second * 3)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go sendMsgToChannel1(c1)
	go sendMsgToChannel2(c2)

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
