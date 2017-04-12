package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "bananas"}

	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range kvs {
		fmt.Println("Key ", key)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println("--------")
	for i, c := range "margo" {
		fmt.Println(i, c)
	}

	fmt.Println("--------")
	for i, c := range "shantul" {
		fmt.Println(i, c)
	}

}
