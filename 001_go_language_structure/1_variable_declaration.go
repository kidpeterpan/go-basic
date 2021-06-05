package main

import "fmt"

func main() {
	var full int = 7
	fmt.Printf("full is %T : %d\n", full,full)

	// go has zero value concept
	var notAssign int
	fmt.Printf("notAssign is %T : %d\n", notAssign, notAssign)

	var noType = 1
	fmt.Printf("noType is %T : %d\n", noType, noType)

	var x,y int = 3,4
	fmt.Printf("x is %T : %d\n", x, x)
	fmt.Printf("y is %T : %d\n", y, y)

	short := 5 // go infer type from the value we assigned
	fmt.Printf("short is %T : %d\n", short, short)

	x,y = y, x // easy value swap
	fmt.Printf("x is %T : %d\n", x, x)
	fmt.Printf("y is %T : %d\n", y, y)
}
