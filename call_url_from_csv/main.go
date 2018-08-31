package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func readFromFile(filePath string) []string {
	f, _ := os.Open(filePath)
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	r.LazyQuotes = true
	var urlsSlice []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for value := range record {
			urlsSlice = append(urlsSlice, record[value])
		}
	}
	return urlsSlice
}

func statusCodes(urls []string) {
	for value := range urls {
		resp, err := http.Get(urls[value])
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode), urls[value])
		}
	}
}

func main() {
	urlsSlice := readFromFile("new_urls.csv")
	statusCodes(urlsSlice)
}
