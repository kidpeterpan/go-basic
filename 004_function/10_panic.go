package main

import "fmt"

func main() {
	// note: น้อยเคสที่จะได้ใช้ panic เพราะมันทำให้โปรแกรมจบการทำงานเลย
	// note: ส่วนมากใช้ในกรณีที่ว่า set ENV config ไม่ครบ เพื่อให้ตอน start server มี config ครบถ้วน
	// เพราะมันไปต่อไม่ได้อยู่ดี
	// common case สำหรับ panic ที่เกิด เกิดจาก เรียกใช้ member จาก struct nil หรือเรียกใช้ arrays หรือ slice ที่ out of range
	defer fmt.Println("is it before or after panic?") // panic จะปล่อยให้ defer ทำก่อน แล้วค่อย panic
	x := make([]int, 10)
	fmt.Println(x[11]) // panic: runtime error: index out of range [11] with length 10
}
