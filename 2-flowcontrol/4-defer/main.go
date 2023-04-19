package main

import "fmt"

func main() {
	defer fmt.Println(":D")
	defer fmt.Println("world")
	fmt.Println("Hello")
}
