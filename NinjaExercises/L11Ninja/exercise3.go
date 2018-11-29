package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("I'm a big ole error %v:", ce.info)
}

func main() {
	e1 := customErr{
		info: "need more caffeine",
	}
	foo(e1)
}

func foo(e error) {
	fmt.Println("Foo ran - ", e, "\n")
}
