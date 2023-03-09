// Prints all prime numbers between a range m : n
// Option: 
// 	[m : n] segmentedSieve(...)
// 	[m : n) segmentedSieve_2(...)
package main

import (
	"fmt"
	"math"
)

func main() {

	var m, n, lines int

	fmt.Scanf("%d", &lines)

	var data [][]int
	
	for lines > 0 {
		fmt.Scanf("%d %d", &m, &n)
		if m > n {
			m,n = n,m
		}
			
		input := []int{m, n}
		data = append(data, input)

		lines--
	}

	for _, data_row := range data {
		if data_row[1] == 3{
			fmt.Println(2)
			fmt.Println(3)
		} else {
			primes := segmentedSieve(data_row[1])
			
			for _, value := range primes {
				if value >= data_row[0] {
					fmt.Println(value)
				}
			}
			fmt.Println()
		}
	}
}

// This function finds all primes smaller or equal than 'n',
// also stores found primes in int slice
func sieveOfEratosthenes(n int)([]int) {
	
	sieve := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			for j := i*i; j <= n; j += i{
				sieve[j] = false
			}  
		}
	}

	var primes []int

	for i := 2; i <= n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
}

// This function finds all primes smaller or equal than 'n'
func segmentedSieve(n int)([]int){

	root := int(math.Sqrt(float64(n)))

	//root primes []int
	r_primes := sieveOfEratosthenes(root)
	
	// Divide the range [0..n] in different segments by 2 index
	// We have chosen segment size as sqrt(n)
	idxLower := root
	idxHigher := 2*root

	// Sieve declaration, size of root (+1 to include reference to higher-bound)
	// also size is equal to [high - lower +1]
	sieve := make([]bool, root+1)

	// While all segments of range are not processed,
	// process one segment at a time
	for idxLower <= n {

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
			// if size is root+1 cycle to <= higher-bound to include it
			// (*) to include range-right-bound (cycle <=) 
			// Including r-bound to check produce for every cycle one iteration more
            // due to the next starting point is equal to the last higher-bound
            for j := start; j <= idxHigher; j += r_primes[i] {
            	sieve[j-idxLower] = false
            }
		}

		// (*) to include range-right-bound (cycle <=)
		for i := idxLower; i <= idxHigher; i++ {
			if sieve[i-idxLower] {	// i%idxLower is also correct, to get the index in current offset
				r_primes = append(r_primes, i)
			}
		}

		idxLower += root
		idxHigher += root
	}

	return r_primes
}

// Segmented version2: This function finds all primes smaller than 'n' -> [0..n)
// With this to get all primes smaller or equal than 'n' -> [0..n] : call func(n+1)
func segmentedSieve_2(n int)([]int){

	root := int(math.Sqrt(float64(n)))

	// root primes []int
	r_primes := sieveOfEratosthenes(root)

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