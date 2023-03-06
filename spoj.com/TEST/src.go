package main

import (
	"fmt"
)

func main(){
	
	var inSlice []int

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