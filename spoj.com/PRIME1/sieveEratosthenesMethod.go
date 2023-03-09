// Sieve of Eratosthenes: given a number n, print all primes smaller or equal than n.

// The sieve of Eratosthenes is one of the most efficient ways to find 
// all primes smaller than n when n is smaller than 10 million or so

package main

import (
	"fmt"
)

func main() {
	

	var n int = 7

	sieveOfEratosthenes(n)

	// Read input from terminal and generate all prime numbers between two given numbers

	// The input begins with the number t of test cases in a single line (t<=10). 
	// In each of the next t lines there are two numbers m and n separated by a space.
	// (1 <= m <= n <= 1000000000, n-m<=100000)
	
	/*
	data := readInput()

	for input_row, _ := range data {
		// get prime less or equal than n
		primes := sieveOfEratosthenes(data[input_row][1])

		// print results
		for i := data[input_row][0]; i <= data[input_row][1]; i++ {
			if primes[i] {
				fmt.Println(i)
			}
		}
		fmt.Println()
	}
	*/
}


func sieveOfEratosthenes(n int) {
	
	// size is n+1 to have aligned index reference
	// sieve{0:1 1:2... 9:10} (n-1:n)
	// 	=> sieve{0:0 1:1... 10:10} (n:n)
	sieve := make([]bool, n+1)

	// Set slice default values true starting from i=2
	// Bool default value is false
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	// Check until square(n) included
	for i := 2; i*i <= n; i++ {
		// If sieve[i] is not changed, then it is a prime
		if sieve[i] {
			// Update all multiples of i greater than or equal to the
			//  square of it.
			// Numbers which are multiple of i and are less than i*i
			//  are already been marked
			for j := i*i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}

	
	fmt.Println("\nSieve bool values: ",  sieve,"\n")
	// custom function return: []bool
	//return sieve

	for i:=2; i<=n; i++{
		if sieve[i] {
			fmt.Println(i)
		}
	}
	
}


func readInput() ([][]int) {

	var lines, m, n int

	fmt.Scanln(&lines)
	
	counter := 0
	var data [][]int = make([][]int, lines)

	for lines > 0 {
		fmt.Scanln(&m, &n)
		// Refactoring
		data[counter] = append(data[counter], m)
		data[counter] = append(data[counter], n)
		lines--
		counter++
	}

	return data
}