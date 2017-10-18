// goRoutesAndChannels
// second round on the video

package main

import (
	"fmt"
	"time"
)

func makeID(idChannel chan int) {
	var id int = 1

	for {
		idChannel <- id
		id += 1
	}

}

func main() {
	start := time.Now()

	// need a channel reciever
	idChannel := make(chan int) // make channel of strings

	// start a go routing
	go makeID(idChannel) // go returns, but emit runs concurrently

	fmt.Printf("%d\n", <-idChannel)

	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}
