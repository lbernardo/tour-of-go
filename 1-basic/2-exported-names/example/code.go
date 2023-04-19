package example

import "fmt"

func IsPublicFn() {
	fmt.Println("Test public fn")
}

var IsPublicVar = "public"

const IsPublicConst = "public"

func isPrivateFn() {

}

var isPrivateVar = "private"

const isPrivateConst = "private"
