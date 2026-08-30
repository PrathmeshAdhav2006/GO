package mylib

import "fmt"

func printHeader(title string) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(title)
	fmt.Println("========================================")
}

func demo_of_not_exported() {
	printHeader("demo_of_not_exported() [unexported]")
	fmt.Println("Calling not_exported() from within the same package (mylib):")
	not_exported()
}

func Demo_of_not_exported() {
	printHeader("Demo_of_not_exported() [exported wrapper]")
	fmt.Println("This wrapper IS exported, even though its name contains 'not_exported'.")
	fmt.Println("Go only checks the first letter — capital D makes this callable externally.")
	fmt.Println("Calling not_exported() internally:")
	not_exported()
}

func DemoOfExported() {
	printHeader("DemoOfExported() [exported wrapper]")
	fmt.Println("Calling Exported() from within the same package:")
	Exported()
}
