package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(p *person, newName string) {
	// p.firstName = newName
	(*p).firstName = newName
}

func main() {
	p1 := person{
		"Mark",
		"Grant",
		36,
	}
	fmt.Println(p1.firstName)
	changeMe(&p1, "Ray")
	fmt.Println((&p1).firstName)
}
