package main

import "fmt"

func main() {
	ex1 := string(rune(0x265e))
	fmt.Println("EX1: integer -> string") // จะได้ string UTF-8 represent
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%x", ex1[i])
	}
	fmt.Println("\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)

	fmt.Println("EX2: []byte -> string")
	ex2 := []byte{0xe2,0x99,0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String) // --> ♞

	fmt.Println("EX3: []rune -> string")
	ex3 := []rune{0xe0,0xb8,0xaa}
	fmt.Println(string(ex3))

	fmt.Println("EX4: string -> []byte")
	ex4 := []byte("Hi-ส")
	fmt.Println(ex4)
	fmt.Println(len(ex4)) // should be 4

	fmt.Println("EX5: string -> []rune")
	ex5 := []rune("Hi-ส")
	fmt.Println(len(ex5)) // should be 6
}
