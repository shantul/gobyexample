package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial value: ", i)

	zeroval(i)
	fmt.Println("Zero Val: ", i)

	zeroptr(&i)
	fmt.Println("Zero Ptr: ", i)

	fmt.Println("Pointer i: ", &i)
}
