package main

import "fmt"

func main() {
	// note: ***HOF คือ function ที่สามารถรับ function และ return function ได้***
	// example with create fibonacci func
	fmt.Println(Fn(7)) // result should be 13
	// add simple func for test fibonacci performance
	testFn := func() func(n int) int {
		fmt.Println("we are testing fibo...")
		return Fn
	}
	fmt.Println(testFn()(7)) // log and get 13
}

// Fn = Fn - 1 + Fn -2 for n > 1
func Fn(n int) int  {
	if n <= 1 {
		return n
	}
	return Fn(n- 1) + Fn(n- 2)
}
