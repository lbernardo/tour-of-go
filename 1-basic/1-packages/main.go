package main

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"math/rand"
)

func main() {
	fmt.Println("Meu número favorito é", rand.Intn(10))
	fmt.Println(aurora.Magenta("Exemplo de pacote"))
}
