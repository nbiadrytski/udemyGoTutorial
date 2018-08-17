package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 { // if i == 2, control will go back to i++
			continue
		}
		fmt.Println(i) // will print only nechetnue nubmers
		if i >= 50 {
			break
		}
	}
}
