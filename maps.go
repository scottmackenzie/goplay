// maps
// associative array, hash, index etc in other programs
package main

import (
	"fmt"
)

func main() {
	// in go we create maps using "Make"
	dayMonths := make(map[string]int) // name of month mapped to day of month
	// any type with a quality can be used string, int, float etc
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in Feb %d\n", dayMonths["Feb"])

	// what if we go for incorrect data?
	fmt.Printf("Days in February %d\n", dayMonths["Feburary"]) // you get 0... (we go 0 due to it being a int)

	// lets check between an error and a correct result using OK syntax
	days, ok := dayMonths["January"]
	if !ok {
		fmt.Printf("Can't get days for January\n\n")
	} else {
		fmt.Printf("%d\n\n", days)
	}

	// lets iterate over a map....

	for month, days := range dayMonths {
		fmt.Printf("%s has %d days\n", month, days) // note it will not come out in order
	}

	// count how many months have 31 days

	has31 := 0
	has30 := 0
	has28 := 0

	for _, days := range dayMonths { // we don't care about months for use _
		if days == 31 {
			has31++
		} else if days == 30 {
			has30++
		} else if days == 28 {
			has28++
		}
	}
	fmt.Printf("Number of months with 31 days: %d\n", has31)
	fmt.Printf("Number of months with 30 days: %d\n", has30)
	fmt.Printf("Number of months with 28 days: %d\n", has28)

	// Lets say we wanted to remove one of the items from the map
	// we use "delete"

	delete(dayMonths, "Feb")
	fmt.Printf("Days in Feb %d\n", dayMonths["Feb"])
	// what if we delete it again
	delete(dayMonths, "Feb") //NOTHING - it doesn't complain
	fmt.Printf("Days in Feb %d\n", dayMonths["Feb"])

	// if we are making a map with a defined set of data we can do it like this...
	dayMonths2 := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	fmt.Printf("Days in Feb %d\n", dayMonths2["Feb"])
	fmt.Printf("Days in February %d\n", dayMonths2["February"])

}
