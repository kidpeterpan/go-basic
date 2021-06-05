package main

import (
	"fmt"
	"go/token"
)

func main() {
	// เราสามารถ explore ที่ token pkg ได้ว่า go มี keywords หรือ operator อะไรใช้เราใช้บ้าง
	fmt.Println(token.ADD) // +
	// กดเข้าไปอ่านใน token pkg เลยเนื่องจาก go เป็น open source อยู่แล้ว

	// ส่วน data type ดูได้จาก builtin/builtin.go


}
