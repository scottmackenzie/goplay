// errors

// Errors are returned from functions and the calling function usually deals with the error
// there are also panic/exceptions functions but they are rearely used - reccomend to not use them

// GO IS BUILT AROUND HANDLING ERRORS RETURNED BY FUNTIONS

package main

import (
	"errors" // an errors package for printer3
	"fmt"
	"os"
	"time"
)

var (
	errorEmptyString = errors.New("I am unable to print an empty string (via errors package)")
)

// Standard error using built in error code
func printer(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	// adds a new line
	return err
}

// If there is an empty string, we should throw an error using Errorf
func printer2(msg string) error {
	if msg == "" {
		return fmt.Errorf("I am unable to print an empty string")
	}
	_, err := fmt.Printf("%s\n", msg)
	// adds a new line
	return err
}

// Same as printer2 but we'll use the "errors" package now
func printer3(msg string) error {
	if msg == "" {
		return errorEmptyString
	}
	_, err := fmt.Printf("%s\n", msg)
	// adds a new line
	return err
}

func main() {
	start := time.Now()

	if err := printer("The brown fox jumps over the lazy dog"); err != nil {
	}

	if err := printer2(""); err != nil {
		fmt.Printf("printer2(): Printer function failed with error: %s\n", err)
	}

	// using errors package
	if err := printer3(""); err != nil {
		fmt.Printf("printer3(): Printer function failed with error: %s\n", err)
	}

	// using errors package but also looking at the kind of error we got back
	if err := printer3(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("You tired to print an empty string!\n")
		} else {
			fmt.Printf("printer3(): Another error occoured: %s\n", err)
		}
	}

	// using panic if something terrible happens
	// using errors package but also looking at the kind of error we got back
	if err := printer3(""); err != nil {
		if err == errorEmptyString {
			panic("You tired to print an empty string!\n")
			os.Exit(1)
		} else {
			fmt.Printf("printer3(): Another error occoured: %s\n", err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
