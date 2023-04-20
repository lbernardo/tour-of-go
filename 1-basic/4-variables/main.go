package main

import "fmt"

var cobol, python, clojure bool

// with initializers
var a int = 1
var b, c int = 2, 3

func main() {
	var i int
	// short variable declarations use :=
	text := "more than words"
	fmt.Println(i, cobol, python, clojure, a, b, c, text)

	x := 5
	fmt.Println("x", x)
	if true {
		x := 10
		fmt.Println("x", x)
	}
	fmt.Println("x", x)
}

/**
Basic types

bool (default: false)

string (default: '')

int int8 int16 int32 int64 (default: 0)
uint uint8 uint16 uint32 uint64 uintptr (default 0)

float32 float64 (default 0)

interface{} <==> any (+go1.8) (default nil)
*/
