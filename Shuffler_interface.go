// Shuffler_interface

// Interfaces are extremely powerful as they allow you
// ... to define behaviour without defining types themselves (e.g. shuffer below).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type shuffer interface {
	// Shuffer is an interface
	// Anything that has these methods defined on it, can be shuffled using shuffler
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffer) { // note we're passing in an interface not a concrete type
	// this allows a sort of object orientation (duck typing)
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i) // pick between the length of s.Len
		s.Swap(i, j)
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// Satisfy the same interface, but with a string slice

type stringSlice []string

func (ss stringSlice) Len() int {
	return len(ss)
}

func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	start := time.Now()

	is := intSlice{1, 2, 3, 4, 5, 6}
	ss := stringSlice{"The", "quick", "brown", "fox", "jumped", "over"}

	shuffle(is)
	shuffle(ss)

	fmt.Printf("%d\n", is)
	fmt.Printf("%s\n", ss)

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
