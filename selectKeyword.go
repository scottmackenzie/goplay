// selectKeyword
// Listen and transmit on multiple channels at one time
// channels are bi directional

package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, done chan bool) { // string channel and done channel
	defer close(wordChannel)
	words := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	var i int = 0 // index into words

	// we will also create a timer here, so that after 3 seconds we need to stop

	t := time.NewTimer(3 * time.Second)

	for { // infinate loop
		select { // new - send a word or terminate
		case wordChannel <- words[i]: // if someone is listening on the wordChannel, they'll get the next word
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-done: // however if it recieves a message on done, stop
			done <- true // you can send a true on the same channel
			return
		case <-t.C: // after 3 second timeout return
			return
		}
	}

}

func main() {
	start := time.Now()

	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh) // launch the go routine with 2 channel inputs

	//for i := 0; i < 100; i++ {
	//	fmt.Printf("%d: %s\n", i, <-wordCh)
	//	}

	for word := range wordCh {
		fmt.Printf("%s\n", word)
	}

	// Removed for the timer example
	//doneCh <- true // stop it once the for loop above stops
	//<-doneCh       // actual value is being discarded until it recieves on the done channel

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
