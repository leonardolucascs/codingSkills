// Prints all prime numbers smaller than n
package main

import (
	"fmt"
	"math"
)

// This function finds all primes smaller than 'n' -> [0..n)
// To get all primes smaller or equal than 'n' -> [0..n] : call func(n+1)
func segmentedSieve(n int)([]int){

	root := int(math.Sqrt(float64(n)))

	// root primes []int
	//r_primes := sieveOfEratosthenes(root)

	// Initialize the sieve
	r_primes := []int{2}


	// Divide the range [0..n) in different segments by 2 index
	// Here I have chosen segment size as sqrt(n)
	idxLower := root
	idxHigher := 2*root

	// Sieve declaration - size of root (+1 to include reference to higher bound)
	// also size is equal to [higher - lower +1]
	sieve := make([]bool, root+1)

	// While all segments of range are not processed,
	// process one segment at a time
	for idxLower < n {

		// if higher goes over n
		if idxHigher > n {
			idxHigher = n
		}

		// Set default values for current segment
		for i := 0; i < len(sieve); i++ {
			sieve[i] = true
		}

		// Sieve the current segment
		for i := 0; i < len(r_primes); i++ {

			// Set start offset
			// Find the minimum number in [l..h] that is a multiple of
			// prime[i], divisible by prime[i]
			start := (idxLower/r_primes[i]) * r_primes[i]
			
			if start < idxLower {
				start += r_primes[i]
			}

			// Mark all multiples of j
            for j := start; j < idxHigher; j += r_primes[i] {
            	sieve[j-idxLower] = false
            }
		}

		for i := idxLower; i < idxHigher; i++ {
			if sieve[i-idxLower] {	// i%idxLower is also correct, to get the index in current offset
				r_primes = append(r_primes, i)
			}
		}

		idxLower += root
		idxHigher += root
	}

	return r_primes
}

func main() {
	n := 7
	primes := segmentedSieve(n)
	fmt.Println(primes)
}
