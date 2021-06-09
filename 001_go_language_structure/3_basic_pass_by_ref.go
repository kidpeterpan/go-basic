package main

import "fmt"

func main() {
	x := 2
	doubleValue(&x)
	fmt.Println("double value of x is:", x)
}

func doubleValue(x *int) {
	*x = *x * 2
}
