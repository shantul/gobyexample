package main

import "fmt"

func main() {

	slice := make([]string, 3)
	fmt.Println("Slice ", slice, " len = ", len(slice), " cap = ", cap(slice))

	slice[0] = "one"
	slice[1] = "two"
	slice[2] = "three"
	fmt.Println("set ", slice)
	fmt.Println("get[2] ", slice[2])

	slice = append(slice, "four", "five")
	fmt.Println("Appended Slice ", slice, " len = ", len(slice), " cap = ", cap(slice))

	newSlice := make([]string, len(slice) + 1)
	copy(newSlice, slice)
	fmt.Println("Copied slice ", newSlice)

	anotherSlice := slice[2:5]
	fmt.Println("Another Slice ", anotherSlice)

	anotherSlice = slice[:4]
	fmt.Println("Another Slice ", anotherSlice)

	anotherSlice = slice[2:]
	fmt.Println("Another Slice ", anotherSlice)

	yetAnotherSlice := []string{"a", "b", "c"}
	fmt.Println("Yet Another Slice ", yetAnotherSlice)

	twoDSlice := make([][]int, 3)
	fmt.Println("Two D Slice ", twoDSlice)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		fmt.Println("Two D Slice ", twoDSlice)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("Two D Slice ", twoDSlice)

	anArray := [3]string{"a", "b", "c"}
	sliceFromArray := anArray[:]
	fmt.Println("Slice from array ", sliceFromArray)
	sliceFromArray = append(sliceFromArray, "d")
	fmt.Println("Slice from array ", sliceFromArray)

	sliceOne := []string{"one", "two", "three", "four", "five"}
	fmt.Println("Slice One ", sliceOne)
	sliceTwo := sliceOne[2:]
	fmt.Println("Slice Two ", sliceTwo)
	sliceTwo[1] = "six"
	fmt.Println("Slice One ", sliceOne)
	fmt.Println("Slice Two ", sliceTwo)
}
