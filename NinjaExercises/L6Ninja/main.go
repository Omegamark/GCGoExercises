package main // package main

import (
	"fmt"
)

// func main() {
	exercise1()
}

func exercise1() {

	fmt.Println(foo())
	fmt.Println(bar())
}
func foo() int {
	return 5
}

func bar() (int, string) {
	return 10, "Woooo!!!!"
}
