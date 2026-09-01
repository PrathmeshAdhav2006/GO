package main

import (
	"fmt"
)

func main() {

	// in go there i no while loop, but we can use for loop to achieve the same functionality
	// for loop type 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	println("-----")

	// for loop type 2
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}
