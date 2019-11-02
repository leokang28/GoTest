package main

import "fmt"

type unit struct{
	Name string
	HP int
}

func main() {
	profile := &unit{
		Name: "rat",
		HP: 100,
	}

	fmt.Printf("'%%v' %v\n", profile)

	fmt.Printf("'%%+v' %+v\n", profile)

	fmt.Printf("'%%#v' %#v\n", profile)

	fmt.Printf("'%%T' %T\n", profile)
}