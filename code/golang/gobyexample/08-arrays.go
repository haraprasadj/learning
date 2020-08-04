package main

import "fmt"

func main() {

	// Type of elements and length are both part of an array's type
	// By default an array is zero-valued
	var a [5]int
	fmt.Println("emp:", a)

	// Set value at an index
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Length of array
	fmt.Println("len:", len(a))

	// Declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can compose
	// types to build multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
