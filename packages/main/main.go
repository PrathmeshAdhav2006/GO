/*
	->In go we cannot have the main function in multiple files.

	->also in same directory we cannot have multiple packages.

	->it is possible that if we create diffrent directories and create package in them
	and can have main function in each of them but we cannot have multiple main functions in same directory.

	->functions can be accessed from other packages if they are exported i.e.
	 the first letter of the function name is capitalized.

	 ->if the first letter of the function name is lowercase,
	 it is not exported and cannot be accessed from other packages.


*/

package main

import (
	"fmt"
	"math"

	"packages/mylib"
)

func printSection(title string) {
	fmt.Println()
	fmt.Println("########################################")
	fmt.Println("# " + title)
	fmt.Println("########################################")
}

func main() {
	printSection("Basic Output")
	fmt.Println("Hello, World!")
	fmt.Println("The square root of 16 is:", math.Sqrt(16))
	greet("abc")

	printSection("mylib: Exported Access")
	mylib.DemoOfExported()

	printSection("mylib: Unexported Access via Exported Wrapper")
	mylib.Demo_of_not_exported()

	printSection("Direct External Access Attempts (would fail if uncommented)")
	fmt.Println("mylib.not_exported()        -> compile error: unexported")
	fmt.Println("mylib.demo_of_not_exported() -> compile error: unexported")
	// mylib.not_exported()
	// mylib.demo_of_not_exported()
}
