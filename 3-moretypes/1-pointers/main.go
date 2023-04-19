package main

import "fmt"

func main() {
	j := 42
	fmt.Println("j", j)
	p := &j
	fmt.Println("p:", *p, "ref:", p)
	*p = 21
	fmt.Println("j", j)
}
