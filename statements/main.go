package main

import (
	"fmt"
	"math"
	"runtime"
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
	} else {
		fmt.Printf("%g >= %g\n", v, lim) //v can be used here because it is in the same block as the if statement
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
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
