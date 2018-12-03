package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() chan int {
			// make 10 goroutines which add 10 numbers each.
			for j := 0; j < 10; j++ {
				c <- j
			}
			return c
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
}
