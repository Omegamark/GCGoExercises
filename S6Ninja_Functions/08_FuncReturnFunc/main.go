package main

import "fmt"

func main() {

	f := returnFunc()
	fmt.Printf("%T\n", f)
	f()
}

func returnFunc() func() {
	return func() {
		fmt.Println(10 * 10)
	}
}
