package main

import "fmt"

func plus(a int, b int) int{
	return a + b
}

func square (a int) int{
	return a * a
}

func main() {
	fmt.Println("Add: ", plus(10, 20))
	fmt.Println("Square: ", square(30))
}