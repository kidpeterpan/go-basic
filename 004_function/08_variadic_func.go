package main

import (
	"fmt"
	"reflect"
)

//variadicFunc use really well with option pattern in go
func variadicFunc(characters ...string){
	fmt.Println(reflect.TypeOf(characters)) // []string
	fmt.Println(characters) // a b c
}

func main() {
	variadicFunc("a","b","c")
	slStr := []string{"a","b","c"}
	variadicFunc(slStr...) // a b c
}
