package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T,%v\n", w, w) // nil, nil
	w = os.Stdout
	fmt.Printf("%T,%v\n", w, w) // *os.File,&{0xc0000520c0}
	b := &bytes.Buffer{}
	fmt.Printf("%T,%v\n", b, b) // *bytes.Buffer,
	b = nil
	w = nil
	fmt.Printf("%T,%v\n", w, w) // nil, nil
	fmt.Printf("b == w: %t\n", b == w)

	// note: *** ต้อง aware ว่า interface มีสองค่า type กับ value
}
