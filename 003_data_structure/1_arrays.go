package main

import (
	"fmt"
	"reflect"
)

func main() {
	// การเก็บข้อมูลแบบเรียงกันลงใน ram
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits)) // 5

	// note - [2]int and [3]int are type different
	twoSlots := [2]int{}
	threeSlots := [3]int{}
	fmt.Println("Two slots:", reflect.TypeOf(twoSlots))
	fmt.Println("Three slots:", reflect.TypeOf(threeSlots))
}
