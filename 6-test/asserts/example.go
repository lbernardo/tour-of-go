package asserts

import "strconv"

func Sum(a, b int) int {
	return a + b
}

func Mult(b, c int) int {
	return b * c
}

func Convert(s string) (int, error) {
	return strconv.Atoi(s)
}
