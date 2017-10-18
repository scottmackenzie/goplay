// goRoutinesandStop
package main

import (
	"fmt"
	"time"
)

func emit(c chan string, done chan bool) {
	t := time.NewTimer(3 * time.Second)

	var words string = "The quick brown fox\n"
	for {
		select {
		case c <- words:
		case <-done:
			fmt.Printf("Told to terminate\n")
			done <- true // send back a close (channel is bi-directional)
			close(done)
			return
		case <-t.C:
			fmt.Printf("Got a terminate\n")
			close(c)
			return
		}
	}
}

func main() {
	start := time.Now()

	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(wordChannel, doneChannel)

	for word := range wordChannel {
		fmt.Printf("%s\n", word)
	}

	doneChannel <- true
	<-doneChannel // wait to confirm close

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
