package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is ", p.name)
}

type Android struct {
	Person
	model string
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{"Ann", 40})

	s := Person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	a := new(Android)
	a.name = "R2D2"
	a.talk()
}