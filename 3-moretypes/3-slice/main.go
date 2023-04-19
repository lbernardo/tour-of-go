package main

import "fmt"

func main() {
	var i []int
	j := []int{1, 2, 3}
	t := j[0:2] // you can use j[:2]

	fmt.Println(i, j, t, len(j), append(j, 4))

	// You can ignore declaration index with '_'
	for index, value := range j {
		fmt.Printf("index: %v value: %v\n", index, value)
	}
}
