package main

import "fmt"

func main() {
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	sliceFruits := fruits[1:4]
	fmt.Println(sliceFruits) // [banana papaya orange]
	fmt.Println(len(sliceFruits)) // 3
	fmt.Println(cap(sliceFruits)) // 4

	refSliceFruits := sliceFruits
	refSliceFruits[0] = "super orange"
	fmt.Println(refSliceFruits[0]) // "super orange"
	fmt.Println(sliceFruits[0]) // "super orange" ถูกเปลี่ยนด้วยเพราะเป็น ref mem address ที่เดียวกัน
}
