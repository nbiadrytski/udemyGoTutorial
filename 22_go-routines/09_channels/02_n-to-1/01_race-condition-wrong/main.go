package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	// 2 goroutines trying to access the same shared var wg

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
