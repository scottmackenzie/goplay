// channelOfChannels
// You can send a channel on a channel

package main

import (
	"fmt"
	"time"
)

func emit(chanChannel chan chan string, done chan bool) { // channel channel channel and done channel
	wordChannel := make(chan string)
	chanChannel <- wordChannel
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

	channelCh := make(chan chan string)
	doneCh := make(chan bool)

	go emit(channelCh, doneCh) // launch the go routine with 2 channel inputs

	wordChan := <-channelCh // pass in the channel to wordChan to use

	for word := range wordChan {
		fmt.Printf("%s\n", word)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
