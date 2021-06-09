package main

import (
	"fmt"
	"github.com/konghiran/001_go_language_structure/prime"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d is prime? -> %t\n", i, prime.IsPrime(i))
	}
}
