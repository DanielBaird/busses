package main

import "fmt"

func main() {

	const Max = 10000
	const PrintCols = 8

	primes := findPrimes(Max)
	for i, p := range primes {
		fmt.Print(p, "\t")
		if i%PrintCols == PrintCols-1 {
			fmt.Print("\n")
		}
	}
}
