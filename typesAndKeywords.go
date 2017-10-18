// typesAndKeywords
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// A user defined type
type webPage struct {
	//A struct in the C programming language (and many derivatives) is a composite data type declaration that defines a physically grouped list of variables to be placed under one name in a block of memory, allowing the different variables to be accessed via a single pointer
	url  string
	body []byte
	err  error
}

func (w *webPage) get() { // receiver is this and it is a pointer (object you're operating on)
	// If you were using a map or a slice (who's type is passed as a refeference)
	// then can safely just use a value (i.e.:  webPage with no *) because the value is passed around
	// If you didn't use a pointer here, a copy would be made of webPage and the w. in the main()
	// section would not see the result from this function

	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOK() bool {
	return w.err == nil
}

func main() {
	start := time.Now()

	w := &webPage{url: "http://www.scott.com"}
	w.get()

	/* you can also set this like this:
	w := &webPage{}				<-- sets the values in w instance of WebPage to 0
	w.url = "http://www.scott.com"
	*/

	/* or like this:
	w := mew(webPage)				<-- sets the values in w instance of WebPage to 0
	w.url = "http://www.scott.com"
	*/

	if w.isOK() {
		fmt.Printf("URL: %s Length %d!\n", w.url, len(w.body))
	} else {
		fmt.Printf("Something went wrong: Error: %s", w.err)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
