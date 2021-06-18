package main

import "fmt"

func main() {
	// defer จะทำการ execute ทีหลังเมื่อโปรแกรมทำงานจบ
	defer fmt.Println("a")
	defer fmt.Println("b") // b จะ execute ก่อน a เพราะตอนเราสั่ง defer มันเป็น stack FILO
	fmt.Println("c")
	fmt.Println("d")
	// note: expression ใน defer จะถูก execute ก่อนนำไปเก็บ
	x := 1
	fn := func(n int) int {
		x := x + n
		return x
	}
	defer fmt.Println(fn(3)) // should be 4
	// ลองมาดูอีกตัวอย่าง
	fn2 := func() func() {
		fmt.Println("1")
		fmt.Println("***")
		return func() {
			fmt.Println("2")
		}
	}
	defer fn2()() // ก่อนโปรแกรมจบ จะทำการ evaluate ก่อน คือ print 1 *** และนำ func print 2 เข้า defer stack
	fmt.Println("=== end ===")

	// note: defer ทำงานหลัง return
	// note: เหมาะมากในการนำมาใช้ close resource
}
