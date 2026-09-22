package secret

// Reveal is a function that returns a string message indicating that it is only importable from within the 'packages' module tree.
func Reveal() string {
	return "This is only importable from within the 'packages' module tree."
}
