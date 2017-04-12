package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Array a ", a)

	a[4] = 100
	fmt.Println("set ", a)
	fmt.Println("get ", a[4])

	fmt.Println("length ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b ", b)

	var twoD [5][5]int
	for i := 0; i < 5; i++ {
		for j :=0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two D Array ", twoD)

	var c = [...]int{1: 10, 2: 20, 5: 50, 200: 200}
	fmt.Println("c len ", len(c))
	fmt.Println(c[50])


}
