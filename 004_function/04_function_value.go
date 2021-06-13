package main

import (
	"errors"
	"fmt"
)

func main() {
	// we want to add two numbers
	x := add(1, 2)
	fmt.Println("x is :", x)
	// we want to sub two numbers
	y := sub(2, 1)
	fmt.Println("y is :", y)
	// จะสังเกตุว่า add() และ sub() เหมือนกันเลย ต่างกันแค่ operator
	// function value จะมาช่วยตรงนี้
	// ลองมาสร้าง apply(int,int,operator)(int, error) กัน -> operator ในที่นี้คือ function signature
	x, _ = apply(1, 2, add)
	y, _ = apply(2, 1, sub)
	fmt.Println("x is :", x)
	fmt.Println("y is :", y)
}

func apply(a int, b int, operator func(a int, b int) int) (int, error) {
	if operator != nil {
		return operator(a, b), nil
	} else {
		return 0, errors.New("unknown operator")
	}
}

func sub(a int, b int) int {
	return a - b
}

func add(a int, b int) int {
	return a + b
}
