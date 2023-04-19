package main

import "fmt"

func main() {
	// Function without return
	example3("Result", 10, 20)
	// Function with args
	fmt.Println(example1(10, 20))
	// Function  with two or more consecutive named function parameters share a type
	fmt.Println(example2(10, 20))
	// Function with multiple results
	fmt.Println(example4("Hello", "World"))
	// Function with named returns
	x, y := example5(10)
	fmt.Printf("X: %v Y: %v", x, y)
}

func example1(x int, y int) int {
	return x + y
}

func example2(x, y int) int {
	return x + y
}

func example3(text string, x, y int) {
	fmt.Println(text, x+y)
}

func example4(a string, b string) (string, string) {
	return a, b
}

func example5(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
