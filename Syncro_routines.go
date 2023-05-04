package main

import (
	"fmt"
)

func main() {
	one := make(chan int)
	two := make(chan int)
	wait := make(chan bool)
	go printPing(one, two)
	go printPong(one, two, wait)
	two <- 1
	<-wait
}

func printPing(one, two chan int) {
	for i := 0; i < 2; i++ {
		<-two
		fmt.Println("Ping")
		one <- 1
	}
	close(one)
}

func printPong(one, two chan int, wait chan bool) {
	for {
		_, ok := <-one
		if ok == false {
			break
		}
		fmt.Println("Pong")

		two <- 1

	}

	wait <- true
}
