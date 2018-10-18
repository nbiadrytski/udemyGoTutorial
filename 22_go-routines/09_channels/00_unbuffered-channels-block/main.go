package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			// give me the next value off the channel
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(time.Second)

	// put 1 int to channel, then code stops
	// until <-ch takes that value off
}
