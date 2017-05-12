package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(from int) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	//synchronous execution

	f(100)

	//concurrent execution
	for i := 0; i < 5; i++ {
		go f(i)
	}

	//go func(msg string) {
	//	fmt.Println(msg)
	//}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}