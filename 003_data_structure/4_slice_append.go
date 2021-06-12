package main

import "fmt"

func main() {
	// สร้าง arrays
	x := [10]int{}
	fmt.Println(x)
	// สร้าง slice a,b จาก arrays

	// ให้ a = index แต่ละตัว ยกกำลัง 2

	// ให้ b = [98 99]

	// a = append(a, b...) // แตก element แต่ละตัวของ b ใส่ต่อท้าย a (append(a, b[0], b[1]))

	// แก้ไข element b ใน a, element ใน b ต้องเปลี่ยนด้วย x ก็จะเปลี่ยนด้วย เพราะมัน share memory กัน


}
