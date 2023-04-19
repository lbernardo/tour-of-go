package main

import "fmt"

func main() {
	sum := 0
	// basic
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional (like 'while')
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// forever
	for {

	}
}
