package main

import "fmt"

func main() {
	c := make(chan int)
	q := gen(c)
	receive(q)
}

func receive(q <-chan int) {
	for v := range q {
		fmt.Println(v)
	}
}

func gen(c chan int) <-chan int {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
