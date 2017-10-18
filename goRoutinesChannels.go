// goRoutinesChannels

// Go routines are seperately running function or thread that runs concurrently with the main program

// Key to how go operates is to pass messages between go routines, which is done via a channel

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func emit(c chan string) { // expects inboud a channel, c, of strings
	words := []string{"The", "Quick", "Brown", "Fox"}

	for _, word := range words {
		c <- word
		//time.Sleep(100 * time.Millisecond) // just so we can see it
	}
	close(c) // closes the channel
}

func emit2(c chan string) { // expects inboud a channel, c, of strings
	words := []string{"The", "Quick", "Brown", "Fox"}

	for _, word := range words {
		c <- word
		//time.Sleep(100 * time.Millisecond) // just so we can see it
	}
	//close(c) // closes the channel
}

func emitRandomNumbers(i chan int) {
	defer close(i)
	count := 0
	for count < 10 {
		i <- rand.Intn(1000)
		count++
	}
}

func makeId(idChannel chan int) {
	var id = 8000000 // our unique ID, starting at 8000000
	for {
		idChannel <- id // send the ID back
		id++            // add 1 to the ID
	}
}

func main() {
	start := time.Now()

	wordChannel := make(chan string) // make a channel capable of passing strings

	// start a go routine
	go emit(wordChannel) //call emit passing in the channel and it runs concurrently with main()

	for word := range wordChannel { // recieve everything on this channel, 1 by 1 until the channel is closed
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	// We don't have to use range to recieve from a channel, we can explicity recieve
	wordChannel = make(chan string) // Make the channel again
	go emit(wordChannel)            // Recall the channel
	word2 := <-wordChannel          // We will get "the"
	fmt.Printf("%s ", word2)
	word2 = <-wordChannel // We will get "quick"
	fmt.Printf("%s ", word2)
	word2 = <-wordChannel // We will get "brown"
	fmt.Printf("%s ", word2)
	word2 = <-wordChannel // We will get "fox"
	fmt.Printf("%s", word2)
	word2 = <-wordChannel // We will get nothing as the channel is closed
	fmt.Printf("\n%s", word2)
	word2, ok := <-wordChannel // We can use ,ok syntax to see what happened
	fmt.Printf("\n%s (%t)", word2, ok)

	// We can start up as many instances as emit as we want

	wordChannel2 := make(chan string) // Make the channel again
	go emit2(wordChannel2)            // Start 1
	go emit2(wordChannel2)            // Start 2
	go emit2(wordChannel2)            // Start 2
	go emit2(wordChannel2)            // Start 2
	go emit2(wordChannel2)            // Start 2
	go emit2(wordChannel2)            // Start 2
	go emit2(wordChannel2)            // Start 2
	// these will all get muddled up as they're going onto the same channel

	// commented out for now due to deadlock example
	/*
		for word2 := range wordChannel2 {
			// deadline means range is waiting for transmission
			fmt.Printf("%s \n", word2)
		}
	*/
	randoms := make(chan int)
	go emitRandomNumbers(randoms)
	fmt.Printf("\n\nPrinting some randoms < 1000\n")
	for n := range randoms {
		fmt.Printf("%d\n", n)
	}

	// Ralistic example is a function to generate a Unique ID
	// It hides all the complexity of ID generation from the main program
	// Also means, if many parts of your program are using a process to generate an ID,
	// 		they are always guarenteed to get an ID that is unqiue & aligned to the ID generation format
	idChannel := make(chan int)
	go makeId(idChannel)
	fmt.Printf("\nID: %d", <-idChannel)
	fmt.Printf("\nID: %d", <-idChannel)
	fmt.Printf("\nID: %d", <-idChannel)

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
