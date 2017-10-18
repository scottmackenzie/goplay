// buffered_Channels
// You don't have to be listening on a channel to transmit to it, up to a size
// if the channel gets full, the writer needs to wait for space
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var running int64 = 0

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d", running)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	atomic.AddInt64(&running, -1)
	fmt.Printf("]")
}

func worker(sema chan bool) {
	<-sema
	work()
	sema <- true
}

func main() {
	sema := make(chan bool, 20)

	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	for x := 0; x < cap(sema); x++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
