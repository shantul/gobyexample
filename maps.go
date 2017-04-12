package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println("Map: ", m)

	k1_value := m["k1"]
	fmt.Println("map[k1] = ", k1_value)

	fmt.Println("Map length: ", len(m))

	delete(m, "k1")

	fmt.Println("Map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	val, key := m["k2"]
	fmt.Println("val, key ", val, key)

	newMap := map[int]string{1 : "a", 2 : "b"}
	fmt.Println("New Map: ", newMap)
}
