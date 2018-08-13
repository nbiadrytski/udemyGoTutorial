package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'
	var h string
	var i int
	var z bool
	var k, l string = "this is k", "this is l"

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)
	fmt.Printf("%v - %T \n", e, e)
	fmt.Printf("%v - %T \n", f, f)
	fmt.Printf("%v - %T \n", g, g)
	fmt.Printf("%v - %T \n", h, h)
	fmt.Printf("%v - %T \n", i, i)
	fmt.Printf("%v - %T \n", z, z)
	fmt.Printf("%v - %T \n", k, k)
	fmt.Printf("%v - %T \n", l, l)
}
