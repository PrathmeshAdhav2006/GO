package mylib

import "fmt"

func not_exported() {
	fmt.Println("→ Inside not_exported() [unexported func in abc.go]")
}

func Exported() {
	fmt.Println("→ Inside Exported() [exported func in abc.go]")
}
