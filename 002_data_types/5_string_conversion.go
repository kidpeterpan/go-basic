package main

import "fmt"

func main() {
	ex1 := string(rune(0x265e))
	fmt.Println("EX1: Integer -> String") // จะได้ string UTF-8 represent
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%x", ex1[i])
	}
	fmt.Println("\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)

	fmt.Println("EX2: []byte -> String")
	ex2 := []byte{0xe2,0x99,0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String) // --> ♞
}
