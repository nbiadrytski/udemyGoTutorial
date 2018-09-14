package main

import "fmt"

func main() {
	n := hashFunction("Go", 12)
	fmt.Println(n)
}

func hashFunction(word string, buckets int) int {
	letter := int(word[0]) // 71
	bucket := letter % 12
	return bucket
}
