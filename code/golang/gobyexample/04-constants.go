package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value
const s string = "constant"

func main() {

	fmt.Println(s)

	// `const` statement can appear anywhere a var statement can
	const n = 500000000

	// Arithmetic is performed with arbitrary precision
	const d = 3e20 / n

	// A numeric constant has no type until it is given one
	fmt.Println(int64(d))

	// A number can be given a type by using in a context that requires one - e.g., `sin` requires float64
	fmt.Println(math.Sin(n))
}
