package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type employee struct {
	id       int
	name     string
	lastName string
}

// receiver attaches this function to person type
func (p person) fullName() string {
	return p.first + " " + p.last
}

func (e employee) fullName() string {
	return e.name + " " + e.lastName
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Mikalai", "Biadrytski", 31}
	e1 := employee{12, "Ted", "B"}

	e1.fullName()

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(e1.fullName())
}
