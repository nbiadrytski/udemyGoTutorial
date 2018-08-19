package main

import "fmt"

func main() {
	fmt.Println(greet("Nick", "B"))

}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
