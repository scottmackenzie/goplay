// scott_leapYear
package main

import (
	"fmt"
)

func leapYear(year int) (result bool) {
	result = year%4 == 0 && year%100 != 0 || year%400 == 0
	//fmt.Printf("%d = %t\n", year, result)
	return result
}

func main() {
	for i := 1799; i <= 2400; i++ {
		if leapYear(i) {
			fmt.Printf("LEAP YEAR: %d\n", i)
		}
	}
}
