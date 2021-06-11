package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	x := "ทดสอบ"
	for index, value := range x {
		fmt.Printf("index: %d, is %c\n",index,value)
	}
	countLenX := utf8.RuneCountInString(x)
	fmt.Println("len of x is:", countLenX)

	for i:=0; i < len(x); {
		r,s := utf8.DecodeRuneInString(x[i:])
		i += s
		fmt.Printf("%c-", r)
	}
}
