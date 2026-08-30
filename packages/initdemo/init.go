package initdemo

import "fmt"

var configLoaded bool

// init runs automatically the moment this package is imported —
// you never call it yourself.
func init() {
	fmt.Println("[initdemo] init() #1 running — loading config...")
	configLoaded = true
}

// You can have MULTIPLE init() functions in the same package/file.
// They run in the order they appear.
func init() {
	fmt.Println("[initdemo] init() #2 running — config loaded:", configLoaded)
}

func Status() {
	fmt.Println("[initdemo] Status() called — configLoaded =", configLoaded)
}
