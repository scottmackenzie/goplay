package main // name of the package and name of file are not related

import (
	"fmt"
)

const (
	// note not using := as it's declared as a const
	messageC     = "The const answer to life is %d\n"
	answerC      = 42
	answerD      = iota * 2 // incrementing value
	answerE                 // set by above (inherited)
	answerResult = "Iota inherited answer:  %d, %d\n"
)

var (
	answerCVar = 42 // You can't add a change to a const, hence Var
)

func main() {

	var message string // type comes at the end
	message = "Hello, World via a var\n"
	message2 := "Hello world via a fast var\n"
	message3 := "The answer to life is %d\n"
	answer := 42
	var pi float64 = 3.14
	pi32 := float32(3.14)
	nine := 9
	nineUnsigned := uint32(9)
	isTrue := true       // a boolean true value
	isFalse := !true     // a boolean false value
	var isSomething bool // in go undeclared variables are given a "0 value" (always initialised)
	b := byte(65)

	fmt.Printf("Hello, World\n") // strings are based on runes (unicode)
	fmt.Printf(message)
	fmt.Printf(message2)
	fmt.Printf(message3, answer)
	fmt.Printf(messageC, answerC)
	// add 1 to a const
	answerCVar += 1
	fmt.Printf(messageC, answerCVar)
	fmt.Printf(answerResult, answerD, answerE)
	fmt.Printf("Pi is: %f\n", pi)
	fmt.Printf("Pi is: %f\n", pi32)
	fmt.Printf("Nine is: %d\n", nine)
	fmt.Printf("Nine Unsigned is: %d\n", nineUnsigned)
	fmt.Printf("Value is: %t and can also be %t\n", isTrue, isFalse)
	fmt.Printf("Undeclared bool is: %t\n", isSomething)
	fmt.Printf("Byte 65 reads as: %x\n", b) // printed in hex

	// exported functions from a package have uppercase
	// private things in a package have a lower case

}
