package main

import (
	"fmt"
)

func approved(student string, result int) {
	if result > 5 {
		fmt.Println(student, "Approved")
	} else {
		fmt.Println(student, "Not approved")
	}
}

func short(x, lim int) int {
	if v := x * 3; v < lim {
		return x
	}
	return lim
}

func main() {
	approved("Paul", 10)
	approved("Jason", 6)
	approved("John", 2)
	approved("Taylor", 8)
	approved("Eric", 7)

	fmt.Println(short(2, 5))
	fmt.Println(short(2, 10))
}
