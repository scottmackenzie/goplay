// byteSlices
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Printf("%s\n", err) // %s often turns something into a string to print it
		os.Exit(1)
	}
	defer f.Close() // at the end close the file

	// read from the file now

	b := make([]byte, 150) // read 100 bytes from the file
	n, err := f.Read(b)    // read into n the number of bytes from the byte slide

	fmt.Printf("%d characters\n%c\n", n, b) // print it out

	// convert it to a string to print it nicely
	fmt.Printf("%d characters\n%s\n", n, string(b)) // print it out

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
