package main

func findPrimes(max int) []int {

	var candidates []int
	var primes []int

	for i := 0; i < max; i++ {
		candidates = append(candidates, i)
	}

	for index, value := range candidates {
		// don't start at 1
		if index < 2 {
			continue
		}

		// is it prime?
		if value != 0 {
			pointer := index
			primes = append(primes, value)
			// strike off all the factors of this prime
			for pointer < max {
				candidates[pointer] = 0
				pointer += value
			}
		}
	}

	return primes
}
