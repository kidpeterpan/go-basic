package main

import "fmt"

func main() {
	// 3 + 4i อะไรประมาณนี้ (จำนวนเชิงซ้อน)
	x := complex(7, 3) // 7 + 3i
	y := complex(1, 7)
	z := x + y
	fmt.Println(z) //8 + 10i
}
