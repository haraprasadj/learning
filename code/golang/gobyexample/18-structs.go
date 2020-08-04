// Structs are typed collections of fields
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Creating a new struct
	fmt.Println(person{"Bob", 20})

	// Naming the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// It is idiomatic to encapsulate new struc creation
	// in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot
	s := person{"Sean", 50}
	fmt.Println(s.name)

	// Dots can be used with struct pointers
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
