package main

import "fmt"

//KB MB GB TB : shifting bits to 10, 20, 30, 40
const (
	_  = iota             //0
	KB = 1 << (iota * 10) //1 * 10
	MB = 1 << (iota * 10) //2 * 10
	GB = 1 << (iota * 10) //3 * 10
	TB = 1 << (iota * 10) //4 * 10
)

func main() {
	fmt.Println("binary \t \t decimal")
	fmt.Printf("%b \t %d \n", KB, KB)
	fmt.Printf("%b \t %d \n", MB, MB)
	fmt.Printf("%b \t %d \n", GB, GB)
	fmt.Printf("%b \t %d \n", TB, TB)
}
