// strings
package main

import (
	"fmt"
)

func main() {

	atoz := "The quick brown fox jumps over the lazy dog\n" // this is our string

	fmt.Printf("%s", atoz)
	fmt.Printf("String for 9 characters: %s\n", atoz[0:9])          // first 9
	fmt.Printf("String for 15 to 19 characters: %s\n", atoz[15:19]) // 15 to 19
	fmt.Printf("String for 15 onwards characters: %s\n", atoz[15:]) // 15 to 19

	// now lets for loop over the string with a count in i
	for i, r := range atoz {
		fmt.Printf("%d = %c\n", i, r)
	}

	// now lets for loop over the string but ignore i as we don't need it
	for _, r := range atoz {
		fmt.Printf("%c-", r)
	}

	// now lets for loop over the string but ignore the data and just count it
	for i := range atoz {
		fmt.Printf("%d- ", i)
	}
	fmt.Printf("\n\n")

	// count the length of the string

	fmt.Printf("The string is %d characters long\n", len(atoz))

	// back quotes example - anything in ` will be taken literally
	atozBack := `"The quick brown "fox" jumps over the lazy dog"`
	fmt.Printf("%s\n", atozBack)

}
