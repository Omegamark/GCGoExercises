package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi my name is %v %v and I am %v years old.", p.first, p.last, p.age)
}

func main() {

	p1 := person{
		"Jack",
		"Flack",
		36,
	}

	p1.speak()

}
