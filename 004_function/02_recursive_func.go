package main

import "fmt"

func main() {
	basicRecursiveCountdown(10) // 10 9 8 7 6 5 4 3 2 1
	fmt.Println(factorial(5))   // 120
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

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) // (5*x) * (4*x) * (3*x) * (2*x) * (1*x) -> x = 1 -> 120
}
