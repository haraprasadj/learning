package main

import "fmt"

// `(int, int)` signature shows that the function returns 2 integers
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Blank identifier _
	_, c := vals()
	fmt.Println(c)
}
