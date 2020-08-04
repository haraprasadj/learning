package main

import "fmt"

func main() {

	// var declares variables
	var a = "initial"
	fmt.Println(a)

	// We can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go infers the type of uninitialized variables
	var d = true
	fmt.Println(d)

	// Uninitialized variables are zero-valued
	var e int
	fmt.Println(e)

	// Shorthand to declare and initialize a variable
	f := "apple"
	fmt.Println(f)
}
