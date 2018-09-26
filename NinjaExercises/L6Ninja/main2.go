// package main

import "fmt"

// func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	defer fmt.Println("I should happen first", n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := bar(ii2)
	fmt.Println("I should happen second", o)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	defer func() {
		fmt.Println("I defered from bar anonymously!!!")
	}()
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
