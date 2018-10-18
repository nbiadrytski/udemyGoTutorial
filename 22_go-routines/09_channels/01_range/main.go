package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// cannot write to channel anymore once it's closed
		// but still can read from channel
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
