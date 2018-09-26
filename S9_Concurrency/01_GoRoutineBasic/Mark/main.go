package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Starting CPUs", runtime.NumCPU())
	fmt.Println("Starting Go Routines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println("I'm r1")
		wg.Done()
	}()

	go func() {
		fmt.Println("I'm r2 yo")
		wg.Done()
	}()

	go func() {
		fmt.Println("I'm r3 Mayne")
		wg.Done()
	}()
	fmt.Println("Middle CPUs", runtime.NumCPU())
	fmt.Println("Middle Go Routines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Ending CPUs", runtime.NumCPU())
	fmt.Println("Ending Go Routines", runtime.NumGoroutine())

}
