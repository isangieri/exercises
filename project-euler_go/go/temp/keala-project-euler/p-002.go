// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import (
	"fmt"
	"time"
)

func fib() (s int) {
	// s is the sum of the even-valued terms
	s = 0
	for first, second := 0, 1; first < 4000000; first, second = second, first+second {
		if first%2 == 0 {
			s += first
		}
	}
	return
}

func main() {
	start := time.Now()

	fmt.Println(fib())

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s", elapsed)
}
