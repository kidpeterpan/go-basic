package main

import "fmt"

func main() {
	fmt.Println("0 is even or odd:", evenOrOdd(0))
	fmt.Println("1 is even or odd:", evenOrOdd(1))
	fmt.Println("2 is even or odd:", evenOrOdd(2))
	fmt.Println("3 is even or odd:", evenOrOdd(3))
	fmt.Println("4 is even or odd:", evenOrOdd(4))
	fmt.Println("5 is even or odd:", evenOrOdd(5))
	fmt.Println("6 is even or odd:", evenOrOdd(6))
	fmt.Println("7 is even or odd:", evenOrOdd(7))
}

func evenOrOdd(num int) string {
	if num % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}