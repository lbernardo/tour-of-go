package main

import "fmt"

type Dog struct {
	Name  string
	Owner Person
}

func (d *Dog) Bark() {
	fmt.Println("Rulf Rulf")
}
