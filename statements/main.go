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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// func pow(x, n, lim float64) float64 {   its same as tha upper func just a diff syntax
// 	v := math.Pow(x, n);                 
// 	if  v < lim {
// 		return v
// 	}
// 	return lim
// }

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
