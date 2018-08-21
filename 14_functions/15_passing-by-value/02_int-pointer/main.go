package main

import "fmt"

func main() {

	age := 44
	fmt.Println("address of age before calling changeMe() ", &age) // 0x82023c080

	changeMe(&age)

	fmt.Println("address of age after calling changeMe() ", &age) //0x82023c080
	fmt.Println("value of age after calling changeMe() ", age)    //24
}

func changeMe(z *int) { // z's type is *int (pointer int)
	fmt.Println("address of z in changeMe() is", z) // 0x82023c080
	fmt.Println("value of z in changeMe() is", *z)  // 44
	*z = 24
	fmt.Println("after dreferencing '*z = 24', z's address is still the same", z) // 0x82023c080
	fmt.Println("z was changed from 44 to 24 by dreferencing '*z = 24'", *z)      // 24
}
