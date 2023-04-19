package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func main() {
	p := Person{
		Name:     "John",
		Lastname: "Wick",
	}
	fmt.Println(p)

	content, _ := json.Marshal(&p)
	fmt.Println(string(content))

	dog := Dog{
		Name: "Daisy",
		Owner: Person{
			Name:     "John",
			Lastname: "Week",
		},
	}
	fmt.Println(dog)

	dog.Bark()

}
