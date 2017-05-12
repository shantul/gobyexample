package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		//fmt.Println("Sending ping...")
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		//fmt.Println("Sending pong...")
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println("...receiving ", <- c)
		//fmt.Println(<- c)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	messages := make(chan string)

	go pinger(messages)
	go ponger(messages)
	go printer(messages)

	var input string
	fmt.Scanln(&input)
}
