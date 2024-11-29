package main

import "fmt"

type Student struct {
	Firstname string
	Lastname  string
}

type Class []Student

func main() {
	// TODO: declare a type for Student (with first and last name)

	mika := Student{"Mika", "Blättler"}
	bob := Student{"Bob", "Miller"}
	hans := Student{"Hans", "Grubenhauer"}
	fmt.Println(mika, bob, hans)

	// TODO: declare a type for Class (consisting of multiple students)
	NewClass := Class{mika, bob, hans}
	SecondClass := Class{bob, mika, hans}
	fmt.Println(NewClass, SecondClass)

	// TODO: declare a map of modules being attended by multiple classes
	// TODO: output everything using fmt.Println()
	modules := map[int][]Class{
		111: {NewClass},
		123: {SecondClass},
		333: {NewClass, SecondClass},
	}
	fmt.Println(modules)
}
