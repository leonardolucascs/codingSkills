package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main(){
	
	var inSlice []int
	
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		n, _ := strconv.Atoi(input)

		if n == 42 {
			break
		} else {
			inSlice = append(inSlice, n)
		}
	}

	var n int

	for {
		fmt.Scan(&n)
		if n == 42 {
			break
		} else {
			inSlice = append(inSlice, n)
		}
	}

	for _, v := range inSlice {
		fmt.Println(v)	
	}
}