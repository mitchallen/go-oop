package main

import (
	"fmt"
)

type Human interface {
	changeFirstName(name string)
}

type Person struct {
	first string
	last  string
	age   int
}

func (person *Person) info() {
	fmt.Printf("%s %s.", person.first, person.last)
	fmt.Printf("\n is %d year(s) old.\n", person.age)
}

func (person *Person) changeFirstName(name string) {
	// person must be * for this to work on original object
	person.first = name
}

func testInferface(h Human, name string) {
	h.changeFirstName(name)
}

// main function
func main() {
	p1 := Person{"John", "Smith", 57}
	p1.info()
	p1.first = "Jane"
	p1.age = 56
	p1.info()
	p1.changeFirstName("Mary")
	p1.info()
	testInferface(&p1, "Zorro")
	p1.info()
}
