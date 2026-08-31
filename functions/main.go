package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func mixed(x string, y int) (int, string) {
	return y + 10, x + " added string"
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(5, 5))
	fmt.Println(mixed("hello", 5))
	a, b := mixed("hello", 5)
	fmt.Println(a)
	fmt.Println(b)

	x, y := split(100)
	fmt.Println(x)
	fmt.Println(y)
}
