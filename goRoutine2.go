// goRoutine2
package main

import (
	"fmt"
	"time"
)

func emit(c chan int) {
	var id int = 1
	for {
		c <- id
		id++
	}
}

func main() {
	start := time.Now()

	intChannel := make(chan int) // make a channel capable of passing strings

	// start a go routine
	go emit(intChannel)

	// receive the info coming down the channel
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)
	fmt.Printf("%d\n", <-intChannel)

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
