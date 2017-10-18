// functionDefinitions
package main

import (
	"fmt"
	//"os"
)

func printer(msg string) {
	fmt.Printf("%s\n", msg)
}

// how about two params into the function and an int
func printer2(msg, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("[%d]: %s : %s\n", repeat, msg, msg2)
		repeat -= 1
	}
}

func printer3(msg string) error {
	// error returns number of bytes and the error.. so
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func printer4(msg string) (string, error) {
	// error returns number of bytes, the converted msg and the error
	msg += "\n"
	_, err := fmt.Printf(msg)
	return msg, err
}

// defer statement
func printer5(msg string) error {
	defer fmt.Printf("\n(line feed)\n") // this runs when the function exists
	// you can have multiple defers, which are fallthrough

	// Returns the error but also uses defer
	_, err := fmt.Printf("%s", msg)
	return err
}

/* defer statement writing to a file (very common)
func printer6(msg b) error {
	f, err := os.Create("/tmp/createdFile.txt") // create a file
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(msg)
	return err
}
*/

func printer7(msg string) (e error) { // decaration of variable e
	// We name the error e
	_, e = fmt.Printf("%s\n", msg)
	return
}

func printer8(format string, msgs ...string) { // multiple strings
	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
	return
}

func main() {
	printer("printer(): Hello World!")                // printer
	printer2("printer2(): Hello World!", "part 2", 4) // printer2
	printer3("printer3(): Hello World!")              // printer3

	appendedMessage, err := printer4("printer4(): Hello World!") //printer4
	if err == nil {
		fmt.Printf("%q\n", appendedMessage) //%q is a 'friendly' printout
		fmt.Printf("%q\n", err)
	}

	printer5("printer5(): Hello World!") // printer5
	//	printer6("printer6(): Hello World to a file!") // printer6
	printer7("printer7(): Hello World!") // printer7

	// lets now take in muliple strings in the function and handle them.
	// Go supports this with ... feature

	printer8("% x\n", "printer7(): Hello World!", "printer7(): Hello World!", "printer7(): Hello World!") // printer7

}
