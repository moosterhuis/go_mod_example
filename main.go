package main

import (
	"fmt"
	"go_mod_example/p1"
	"go_mod_example/p2"

	"github.com/GoesToEleven/puppy"
)

func main() {
	p1.Hello1()
	p2.Hello2()

	fmt.Println(p2.P2)
	puppy.From11()
}
