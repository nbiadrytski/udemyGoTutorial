package main

import (
	"fmt"
)

func main() {

	// 1. program enters on main
	// 2. create 2 channels
	// 3. launch 3 goroutines
	// 4. meanvhile range c is waiting to receive
	// 5. once goroutine#1 loop finished, it sets <-done = true
	// 6. once goroutine#2 loop finished, it sets <-done = true
	// 7. both <-done received true and then channel c is closed
	// 8. range c sees that channel c is closed, so it receives values and assigns them to n and prints them

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
