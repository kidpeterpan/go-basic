package main

import "fmt"

func main() {
	x := 7
	fmt.Println("value of x: ", x)
	fmt.Println("address of x: ", &x)

	var examplePointer *int
	examplePointer = &x // address of examplePointer point to x address
	fmt.Println("address of examplePointer:", examplePointer)

	// so when we change the value of expamplePointer value of x should be change
	*examplePointer = 14
	fmt.Println("value of examplePointer:", *examplePointer)
	fmt.Println("value of x:", x)

	// normal value copy
	y := x
	fmt.Println("y == x", y == x)
	// value of x should not be change
	y = 15
	fmt.Println("y == x", y == x)
	fmt.Println("value of y:", y)
	fmt.Println("value of x:", x)

	// this below code should be false because it always return new address from the getPointer function
	fmt.Println("compare function return pointer getPointer == getPointer:", getPointer() == getPointer())
	// go provide new() for allocate memory it work at the same getPointer() we wrote
	fmt.Println("compare function return pointer new(int) == new(int):", new(int) == new(int))
}

func getPointer() *int {
	x := 5
	return &x
}
