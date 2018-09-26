package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I'm a human person")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Mark",
		36,
	}
	saySomething(&p1)
}
