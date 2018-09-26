package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementor := 0
	gs := 100

	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			fmt.Println(incrementor)
			v := incrementor
			// runtime.Gosched()
			v++
			incrementor = v
			fmt.Println(incrementor)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value", incrementor)
}
