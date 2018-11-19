package main

import (
	"fmt"
	"sync"
)

var count int64

func main() {
	a := incrementor(1)
	b := incrementor(2)
	c := incrementor(3)

	for n := range merge(a, b, c) {
		fmt.Println(n)
	}
}

func incrementor(s int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			s = s + i
			fmt.Println("Process: ", s, " printing:", i)
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
