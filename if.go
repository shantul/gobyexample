package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if i :=9; i < 0 {
		fmt.Println(i, " is negative")
	} else if i < 10 {
		fmt.Println(i, " is a single digit number")
	} else {
		fmt.Println(i, " is a multiple digit number")
	}
}
