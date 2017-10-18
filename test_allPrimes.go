// test_allPrimes
// Challenge - Write a program that prints all prime numbers. (Note: if your programming language does not support arbitrary size numbers, printing all primes up to the largest number you can easily represent is fine too.)
package main

import (
	"fmt"
	//"math"
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

	// Only primes less than or equal to N will be generated
	/*	const N = 5000

		var x, y, n int
		nsqrt := math.Sqrt(N)

		is_prime := [N]bool{}

		for x = 1; float64(x) <= nsqrt; x++ {
			for y = 1; float64(y) <= nsqrt; y++ {
				n = 4*(x*x) + y*y
				if n <= N && (n%12 == 1 || n%12 == 5) {
					is_prime[n] = !is_prime[n]
				}
				n = 3*(x*x) + y*y
				if n <= N && n%12 == 7 {
					is_prime[n] = !is_prime[n]
				}
				n = 3*(x*x) - y*y
				if x > y && n <= N && n%12 == 11 {
					is_prime[n] = !is_prime[n]
				}
			}
		}

		for n = 5; float64(n) <= nsqrt; n++ {
			if is_prime[n] {
				for y = n * n; y < N; y += n * n {
					is_prime[y] = false
				}
			}
		}

		is_prime[2] = true
		is_prime[3] = true

		primes := make([]int, 0, 1270606)
		for x = 0; x < len(is_prime)-1; x++ {
			if is_prime[x] {
				primes = append(primes, x)
			}
		}

		// primes is now a slice that contains all primes numbers up to N
		// so let's print them
		for _, x := range primes {
			fmt.Println(x)
		}
	*/
	elapsed := time.Since(start)
	fmt.Printf("Calc took %s\n", elapsed)
}
