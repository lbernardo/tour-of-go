package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
}

func main() {
	p := map[string]Person{
		"john wick 4": {
			Name:     "John",
			Lastname: "Wick",
		},
		"spider-man": {
			Name:     "Peter",
			Lastname: "Parker",
		},
	}

	for name, person := range p {
		fmt.Printf("Movie: %v -> %v\n", name, person.Name)
	}
}
