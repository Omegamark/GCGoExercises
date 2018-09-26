package main

import "fmt"

func main() {
	g := func(xi []int) int {
		n := 0

		for _, v := range xi {
			n += v
		}
		return n
	}

	x := callback(g, []int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(x)

}

func callback(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
