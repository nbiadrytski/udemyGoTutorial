package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	fmt.Println(p1.First) // empty
	fmt.Println(p1.Last)  // empty
	fmt.Println(p1.Age)   // 0

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`) // converting a string of text to slice of bytes
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.First)   // James
	fmt.Println(p1.Last)    // Bond
	fmt.Println(p1.Age)     // 20
	fmt.Printf("%T \n", p1) // main.person
}
