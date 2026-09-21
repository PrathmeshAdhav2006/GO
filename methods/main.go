package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// The Abs method calculates the absolute value (magnitude) of a Vertex
// instance using the Pythagorean theorem. It takes a Vertex as a receiver and
// returns the square root of the sum of the squares of its X and Y coordinates.

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
