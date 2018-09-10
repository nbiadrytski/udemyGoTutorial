package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

func checkResponseBody(urls []string) {
	for value := range urls {
		resp, err := http.Get("https://media-utils-uat.mtvnservices.com/services/MediaGenerator/" + urls[value] + "?ep=dcc3faa5&device=iPad7,2")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bodyData, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			log.Fatal(err)
		}

		checkedString := strings.Contains(string(bodyData), "http://mtvplus-vh.akamaihd.net")
		if !checkedString {
			fmt.Println(urls[value])
			fmt.Println(string(bodyData))
		}
	}
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

func elapsed(app string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s was executed in %v\n", app, time.Since(start))
	}
}

func main() {
	defer elapsed("app")()
	//urlsSlice := readFromFile("new_urls.csv")
	mgidSlice := readFromFile("mtvplus.csv")
	//statusCodes(urlsSlice)
	checkResponseBody(mgidSlice)
	//time.Sleep(time.Second * 2)
}
