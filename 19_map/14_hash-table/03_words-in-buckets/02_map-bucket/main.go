package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()
	scanner.Split(bufio.ScanWords)

	// Create map with a key of int
	// and a value of another map
	// with a key of string, which will be the word
	// and a value of int, which will be the number of times the word occurs
	buckets := make(map[int]map[string]int)

	// Create slices to hold words words
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n][word]++
	}

	// number of buckets
	fmt.Println(len(buckets))

	// print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// Print words in a bucket (how many times each word occurs)
	//for k, v := range buckets[6] {
	//	fmt.Println(v, " \t- ", k)
	//}
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
