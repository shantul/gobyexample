package main

import "fmt"

func main() {
	sliceOne := []string{"one", "two", "three"}
	fmt.Println("Slice One: ", sliceOne)
	sliceTwo := make([]string, len(sliceOne))
	fmt.Println("Slice Two: ", sliceTwo)
	fmt.Println("------------")

	copy(sliceTwo, sliceOne)
	fmt.Println("Slice One: ", sliceOne)
	fmt.Println("Slice Two: ", sliceTwo)
	fmt.Println("------------")

	sliceThree := sliceOne[1:]
	fmt.Println("Slice One: ", sliceOne)
	fmt.Println("Slice Two: ", sliceTwo)
	fmt.Println("Slice Three: ", sliceThree)
	fmt.Println("------------")

	sliceThree[1] = "ten"
	fmt.Println("Slice One: ", sliceOne)
	fmt.Println("Slice Two: ", sliceTwo)
	fmt.Println("Slice Three: ", sliceThree)
	fmt.Println("------------")

	copy(sliceThree, sliceTwo)
	fmt.Println("Slice One: ", sliceOne)
	fmt.Println("Slice Two: ", sliceTwo)
	fmt.Println("Slice Three: ", sliceThree)
}
