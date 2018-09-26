package main

import (
	"fmt"
	"math"
)

var f int = 7
var g func() float64

var h func() = func() {
	fmt.Println("funcy")
}

func main() {
	x := func() float64 {
		return math.Pi * math.Pi
	}
	h()
	fmt.Println(x())
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
	g = x
	fmt.Println(g())
}
