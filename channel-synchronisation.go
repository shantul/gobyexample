package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
	done <- true
}


func main() {
	ch1 := make(chan bool, 1)
	go worker(ch1)
	//go worker(ch1)
	//go worker(ch1)
	//go worker(ch1)
	//go worker(ch1)

	//ch2 := make(chan bool, 1)
	//go worker(ch2)


	<-ch1
	//<-ch1
	//<-ch1
	//<-ch1
	//<-ch1
	//<-ch1
	//<-ch2
}