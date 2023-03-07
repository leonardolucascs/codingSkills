
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slOutput := leggiInput()

	for row, _ := range slOutput {
		for i := slOutput[row][0]; i <= slOutput[row][1]; i++ {
			if isPrime(i) && i != 1 {
				fmt.Println(i)
			}
		}
		fmt.Println()
	}
}

// Determina i divisori partendo da 2 
// per verificare se n è primo
//  0. metodo naive: cicla tutti gli n numeri
//	1. ciclo fino a n/2

// 	MIGLIORI OTTIMIZZAZIONI
// 	2. ciclo fino a sqrt(n) o i*i <= n
//	3. sieve of Erastothenes
//	4. segmented sieve

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Lettura[complessa] da stdin(terminale o file) si può fare più semplice la lettura
// Implementazione: salvataggio dati letti in slice multidimensionale

func leggiInput() ([][]int) {

	var lines int

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		tmp := scanner.Text()
		lines, _ = strconv.Atoi(tmp)
	}

	var slOutput [][]int = make([][]int, lines)

	// Refactoring: eliminare utilizzo contatore
	counter := 0

	for lines > 0 {
		if scanner.Scan() {
			// Refactoring: lettura diretta input senza passare da conversione
			input := strings.Split(scanner.Text(), " ")
			m, _ := strconv.Atoi(input[0])
			n, _ := strconv.Atoi(input[1])
			// Refactoring: eliminare utilizzo contatore, 
			//	utilizzo []int ausiliaria da salvare dentro [][]int
			slOutput[counter] = append(slOutput[counter], m)
			slOutput[counter] = append(slOutput[counter], n)
			lines--
			counter++
		}
	}

	return slOutput
}