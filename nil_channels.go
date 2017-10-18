// nil_channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(10 * time.Second) // changed 3 to 10
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-t.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	// added later
	t := time.NewTimer(2 * time.Second)
	restarter := time.NewTimer(5 * time.Second)
	// added later
	saved := ch
	for {
		select {
		case ch <- rand.Intn(42):
		case <-t.C:
			ch = nil
		case <-restarter.C:
			ch = saved
		}
	}
}

func main() {
	start := time.Now()

	ch := make(chan int)

	go reader(ch)
	go writer(ch)

	time.Sleep(30 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
