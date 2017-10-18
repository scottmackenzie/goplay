// test_allPrimes
// Challenge - Write a program that prints all prime numbers. (Note: if your programming language does not support arbitrary size numbers, printing all primes up to the largest number you can easily represent is fine too.)
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 2; i <= 5000; i++ {

		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 {

				isPrime = false
			}
		}

		if isPrime == true {

			fmt.Println(i)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Calc took %s\n", elapsed)
}
