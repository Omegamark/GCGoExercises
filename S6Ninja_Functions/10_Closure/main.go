package main

import "fmt"

func main() {
	fmt.Println("I'm the result of enclosed var x:", enclosedVariable([]int{1, 2, 3, 4, 5, 6}))
	x := "I'm the outer var x"
	fmt.Println(x)

	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func enclosedVariable(n []int) int {
	x := 100
	y := 0
	for _, v := range n {
		y += v
	}
	z := x + y
	return z
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
