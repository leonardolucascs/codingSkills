package main

import "fmt"

func main(){
	var lines, m, n int

	fmt.Scan(&lines)

	for lines > 0 {
		//fmt.Scan(&m)
		//fmt.Scan(&n)
		fmt.Scan(&m, &n)

		for i := m; i <= n; i++ {
			if isPrime(i) {
				fmt.Println(i)
			}
		}
		lines--
		fmt.Println()
	}
}

func isPrime(n int) bool {

	if n < 2{; return false };
	if n == 2{; return true };
	if n % 2 == 0{; return false };

	for i := 3; i*i <= n; i += 2 {
		if n % i == 0{; return false};
	}

	return true
}