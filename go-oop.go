package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func (person Person) info() {
	fmt.Printf("%s %s.", person.first, person.last)
	fmt.Printf("\n is %d year(s) old.\n", person.age)
}

// main function
func main() {
	p1 := Person{"John", "Smith", 57}
	p1.info()
	p1.first = "Jane"
	p1.age = 56
	p1.info()
}
