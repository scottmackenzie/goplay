// test_computes

// Challenge:  write a program that computes:  4\cdot \sum_{k=1}^{10^6} \frac{(-1)^{k+1}}{2k-1} = 4\cdot(1-1/3+1/5-1/7+1/9-1/11\ldots).
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Hello World!")

	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}
