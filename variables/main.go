package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool

var x int = 42 //we can do this

//y ,z:= 42, "shiva"   we can do this because we are assigning values to variables inside a function body

// python = true
// java = true  we cannot do this because we are not allowed to assign values to variables outside of a function body

func main() {
	var i int
	c = true

	y, z := 42, "shiva" // we can do this because we are assigning values to variables inside a function body

	fmt.Println(i, c, python, java)

	fmt.Println(y, z) //we need to use y and z because they are local variables we need to use them to avoid err
	//this is not case with package-level variables because they are exempt from the "declared-but-unused" rule

	//Statements (assignment, :=) → only inside functions.
	//Declared-but-unused locals → compile error, only inside functions (package-level declarations are exempt).

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
