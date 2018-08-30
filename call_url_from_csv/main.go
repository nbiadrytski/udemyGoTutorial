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
	r := csv.NewReader(bufio.NewReader(f))
	var urlsSlice []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
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
	urlsSlice := readFromFile("/Users/mikalai_biadrytski/Desktop/new_urls.csv")
	statusCodes(urlsSlice)
}