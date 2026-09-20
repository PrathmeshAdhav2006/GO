package main

import "fmt"

// This program demonstrates the use of buffered channels in Go. A buffered channel allows you to send
// a limited number of values into the channel without requiring a corresponding receiver for each send operation.
// In this example, we create a buffered channel with a capacity of 2, send two integers into it, and then
// receive and print those integers.
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
