package main

import (
	"fmt"
	"math"

	"packages/initdemo"
	"packages/internal/secret"
	m "packages/mylib"
	"packages/structdemo"
	"packages/vars"
)

/*
	-> In Go we cannot have the main function in multiple files
	   within the same package/directory.

	-> Also, in the same directory we cannot have multiple packages —
	   every .go file in a folder must declare the same package name.

	-> It is possible that if we create different directories and create
	   a package in each, every one of them can have its own main function —
	   but we cannot have multiple main functions in the same directory.

	-> Functions can be accessed from other packages if they are exported,
	   i.e. the first letter of the function name is capitalized.

	-> If the first letter of the function name is lowercase, it is not
	   exported and cannot be accessed from other packages.
*/

func printSection(title string) {
	fmt.Println()
	fmt.Println("########################################")
	fmt.Println("# " + title)
	fmt.Println("########################################")
}

func main() {

	/*
		-> init() is a special function every package can define.
		-> It runs AUTOMATICALLY the moment the package is imported —
		   you never call it yourself, and it always runs BEFORE main().
		-> A package/file can have MULTIPLE init() functions; they run
		   in the order they appear.
		-> That's why initdemo's init() output already ran before this
		   line even executes — proof it fired at import time, not here.
	*/
	printSection("init() already ran automatically at import time")
	initdemo.Status()

	printSection("Basic Output")
	fmt.Println("Hello, World!")
	fmt.Println("The square root of 16 is:", math.Sqrt(16))
	greet("abc")

	/*
		-> Import aliasing: "m \"packages/mylib\"" renames the package
		   to "m" for this file only. You then MUST use "m.XXX", not
		   "mylib.XXX" — the original name is no longer valid here.
		-> Useful for shortening long names or avoiding clashes when
		   two imported packages would otherwise share the same name.
	*/
	printSection("Import Aliasing (mylib imported as 'm')")
	m.DemoOfExported()

	printSection("Unexported Access via an Exported Wrapper")
	m.Demo_of_not_exported()

	/*
		-> Struct field exporting: capitalization rules apply to struct
		   FIELDS too, not just functions — even if the struct itself
		   is exported, each field needs its own capital letter to be
		   visible outside the package.
		-> Common pattern: keep a field unexported and expose it only
		   through an exported getter method (see structdemo.User.Age()).
	*/
	printSection("Struct Field Exporting")
	structdemo.Demo()

	/*
		-> Internal packages: any package living under a folder literally
		   named "internal/" can only be imported by code inside that
		   same module tree.
		-> secret.Reveal() works here because main.go is part of the
		   "packages" module — an external module could NOT import
		   "packages/internal/secret" even if it imported "packages".
	*/
	printSection("Internal Package Access")
	fmt.Println(secret.Reveal())

	/*
		-> Exported vs unexported applies equally to package-level
		   variables and constants, not just functions/structs.
		-> vars.ExportedVar / vars.MaxSize are visible here because
		   they start with a capital letter.
		-> vars.unexportedVar / vars.minSize are NOT visible here —
		   they can only be read via an exported function inside
		   the vars package itself (see vars.ShowInternal()).
	*/
	printSection("Exported vs Unexported Vars/Consts")
	fmt.Println("vars.ExportedVar:", vars.ExportedVar)
	fmt.Println("vars.MaxSize:", vars.MaxSize)
	vars.ShowInternal()

	/*
		-> These calls are commented out on purpose — uncommenting any
		   of them causes a COMPILE ERROR, because each refers to an
		   unexported identifier from outside its own package.
	*/
	printSection("Things that would FAIL if uncommented")
	fmt.Println("m.not_exported()      -> unexported func")
	fmt.Println("vars.unexportedVar    -> unexported var")
	fmt.Println("vars.minSize          -> unexported const")
	// m.not_exported()
	// fmt.Println(vars.unexportedVar)
	// fmt.Println(vars.minSize)
}
