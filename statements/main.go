package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 { // the expression need not be surrounded by parentheses ( ) but the braces { } are required.
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
