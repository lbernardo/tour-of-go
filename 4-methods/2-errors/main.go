package main

import (
	"fmt"
	"strconv"
)

func convert(value string) {
	i, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}

func main() {
	convert("42")
	convert("Number 42")
}
