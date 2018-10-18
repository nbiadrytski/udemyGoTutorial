package main

import (
	"fmt"
)

func main() {

	// 1. program enters on main
	// 2. create 2 channels
	// 3. looping n time (10)
	// 4. each time we loop, we launch anonumous func (which is a goroutine)
	// 5. anonumous func loops 10 times. once it's done, it writes true to all 10 done channels
	// 6. 2nd goroutine reads from done channel 10 times (now all done are true)
	// 7. channel c is closed
	// 8. range c sees that channel c is closed, so it reads values and assigns them to n and prints them

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		// self-executing anonumous func
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
