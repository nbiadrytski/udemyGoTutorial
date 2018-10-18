package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. create channel
	// 2. create wg var
	// 3. then 3 goroutines get launched
	// 4. waiting to receive on channel c (for n := range c)
	// 5. when 2 first anonymous funcs done looping 10 times, the 3rd func waits until they are done and closes channel
	// 6. when channel is closed, range will read and assign values from channel to n
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
