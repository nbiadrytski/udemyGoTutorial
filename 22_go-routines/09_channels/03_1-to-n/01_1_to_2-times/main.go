package main

import (
	"fmt"
)

func main() {

	// 1. program enters on main
	// 2. create 2 channels
	// 3. launching all 3 goroutines
	// 4. 1st goroutine loops 1000 and puts the values into c channel
	// 5. 2 other go routines reading values

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done
}
