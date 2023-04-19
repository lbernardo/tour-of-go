package main

import (
	"fmt"
	"github.com/lbernardo/tour-of-go/1-basic/2-exported-names/example"
)

func main() {
	fmt.Println("Constant", example.IsPublicConst)
	fmt.Println("Variable", example.IsPublicVar)
	example.IsPublicFn()
}
