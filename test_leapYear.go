// test_leapYear

// Challenge - Write a program that prints the next 20 leap years
/* Leap year

1. The year is evenly divisible by 4;
2. If the year can be evenly divided by 100, it is NOT a leap year, unless;
3. The year is also evenly divisible by 400. Then it is a leap year.

*/

package main

import (
	"fmt"
	"time"
)

func leapYear(year int) (result bool) {
	result = year%4 == 0 && year%100 != 0 || year%400 == 0
	//fmt.Printf("%d = %t\n", year, result)
	return result
}

func main() {

	start := time.Now()

	year, _, _ := time.Now().Date() // get the year, ignore month, day
	var stop bool
	totalLeapYears := []int{} // slice where we store our leap years

	for !stop {
		if leapYear(year) {
			totalLeapYears = append(totalLeapYears, year)
		}
		year++
		if len(totalLeapYears) == 20 {
			stop = true
		}
	}
	//we have reached 20 items in the slice (i.e. 20 leap years found)
	for i, year := range totalLeapYears {
		fmt.Printf("%d: %d\n", i, year)
	}

	elapsed := time.Since(start)
	fmt.Printf("Calc took %s\n", elapsed)
}
