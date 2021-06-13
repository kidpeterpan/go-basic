package main

import "fmt"

func main() {
	basicRecursiveCountdown(10) // 10 9 8 7 6 5 4 3 2 1
}

// note: recursive ประกอบด้วยสองส่วนหลักๆ
// 1. function เรียกตัวมันเองซ้ำๆ
// 2. terminal condition เมื่อเข้าเงื่อนไขแล้วจะหยุดการทำงาน
func basicRecursiveCountdown(x int) {
	if x == 0 {
		fmt.Printf("\n")
		return
	}
	fmt.Print(x, " ")
	basicRecursiveCountdown(x - 1)
}