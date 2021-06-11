package main

import "fmt"

func main() {
	ex1 := string(rune(0x265e))
	fmt.Println("EX1: Integer -> String") // จะได้ UTF-8
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%x", ex1[i])
	}
	fmt.Println("\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)
}
