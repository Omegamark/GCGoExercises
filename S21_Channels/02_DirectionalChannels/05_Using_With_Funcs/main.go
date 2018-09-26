package main

import "fmt"

func main() {

	c := make(chan int)
	// send
	go foo(c)
	// receive - Both routines are sent and run in parallel with each other and func main. The result is that the program exits before bar() has an opportunity to Println.
	// go bar(c)
	// receive - By removing `go` from the workflow, the program allows the channel to block until it has received the value it expects.
	bar(c)
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
