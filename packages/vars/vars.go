package vars

import "fmt"

var ExportedVar = "I'm a visible package-level variable"
var unexportedVar = "I'm hidden outside this package"

const MaxSize = 100 // exported constant
const minSize = 1   // unexported constant

func ShowInternal() {
	fmt.Println("unexportedVar:", unexportedVar)
	fmt.Println("minSize:", minSize)
}
