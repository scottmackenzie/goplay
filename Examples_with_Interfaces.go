// Examples_with_Interfaces
package main

import (
	"fmt"
	"time"
)

// define an animal interface
type Animal interface {
	speak() string
}

// If you understand that an interface value is two words wide and it contains a pointer
// to the underlying data, that’s typically enough to avoid common pitfalls.

// Now get some animals to use that interface

type Dog struct{}
type Cat struct{}
type Llama struct{}
type JavaProgrammer struct{}

func (d Dog) speak() string { // Dog animal which returns a string
	return "Woof!"
}

func (c Cat) speak() string { // Cat animal which returns a string
	return "Meow!"
}

func (l Llama) speak() string { // Llama animal which returns a string
	return "??? something!"
}

func (j JavaProgrammer) speak() string { // JavaProgrammer animal which returns a string
	return "Design patterns!"
}

func main() {
	start := time.Now()

	// Within the animals slice, each element is of Animal type,
	// ... but our different values have different underlying types.
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		fmt.Printf("%s\n", animal.speak())
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
