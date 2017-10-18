// multipleReadersAndWriters
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// use go routines to download and check page sizes

func getPage(url string) (int, error) {

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	// use ReadAll

	defer resp.Body.Close() // need to close the body after
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil // return length and error

}

func worker(urlCh chan string, sizeChannel chan string, id int) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err != nil {
			sizeChannel <- fmt.Sprintf("ERROR: on URL %s, %s\n", url, err)
		} else {
			sizeChannel <- fmt.Sprintf("(%d): %s is length %d\n", id, url, length)
		}
	}
}

func main() {
	start := time.Now()

	urls := []string{"http://www.google.com", "http://www.yahoo.co.uk", "http://www.bing.com", "http://news.bbc.co.uk", "http://www.edapt.org.uk"}

	urlCh := make(chan string)
	sizeChannel := make(chan string)

	for i := 0; i < 5; i++ {
		go worker(urlCh, sizeChannel, i)
	}

	for _, url := range urls {
		urlCh <- url
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf(<-sizeChannel)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
