package main

import "fmt"

func main() {
	fmt.Println(greet("Nick", "B"))

}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
