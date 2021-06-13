package main

import "fmt"

var x int

func fn() int {
	x = x + 1
	return x
}

func anFn() func() int {
	y := 0 // closure
	return func() int {
		y = y + 1
		return y
	}
}

func main() {
	// เราอยากได้ func ที่เมื่อทุกครั้งที่เรียก return ค่า เพิ่ม 1
	// เราอยากได้ผลลัพธ์ดังนี้ 1, 2, 3, 1, 2, 3
	// concept ของ anonymous func ช่วยเราได้

	// ex1. ไม่ใช้ anonymous func
	fmt.Println(fn()) // 1
	fmt.Println(fn()) // 2
	fmt.Println(fn()) // 3
	// ความเจ็บปวดคือถ้าจะทำอีกรอบ โดยไม่ set x = 0 เราทำไม่ได้
	x = 0
	fmt.Println(fn()) // 1
	fmt.Println(fn()) // 2
	fmt.Println(fn()) // 3

	// ex2. ใช้ anonymous
	f := anFn()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	f2 := anFn()
	fmt.Println(f2()) // 1
	fmt.Println(f2()) // 2
	fmt.Println(f2()) // 3
	// note: เราใช้ closure แทน ไม่ต้องมี global variable เลย
}
