package main

import "fmt"

func main() {
	// note: ส่วนมาเราใช้ recover เพื่อ rethrow panic ให้มีข้อความชัดเจนมากขึ้น
	// note: recovery จะทำงานเฉพาะใน defer
	p := recover()
	fmt.Println(p) // nil

	printA()         // a , recover from: for no reason
	fmt.Println("b") // b
	fmt.Println("c") // c
}

func printA() {
	defer func() {
		if p := recover(); p != nil { // ตอนที่ panic ให้มาทำอันนี้แทน ไม่ใช่หยุด program
			fmt.Println("printA: recover from:", p)
		}
	}()

	fmt.Println("a")
	panic("for no reason")
	fmt.Println("b") // not execute but main func not stop
}
