package main

import "fmt"

type Mammal interface {
	MyNameIs() string
}

type Human struct {
	Name     string
	Lastname string
}

func (h *Human) MyNameIs() string {
	return fmt.Sprintf("%v %v", h.Name, h.Lastname)
}

func main() {
	p := &Human{
		Name:     "John",
		Lastname: "Wick",
	}
	print(p)
}

func print(person Mammal) {
	fmt.Println(person.MyNameIs())
}
