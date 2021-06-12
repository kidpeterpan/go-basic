package main

import "fmt"

func main() {
	// สร้าง arrays
	x := [10]int{}
	fmt.Println(x)
	// สร้าง slice a,b จาก arrays
	a := x[:5]
	b := x[5:7]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("\n")
	// ให้ a = index แต่ละตัว ยกกำลัง 2
	for index, _ := range a {
		a[index] = index * index
	}
	// ให้ b = [98 99]
	b[0] = 98
	b[1] = 99
	fmt.Println(x) // ค่า x เปลี่ยนด้วย เพราะ a,b อ้างถึง address ของ x
	fmt.Println(a)
	fmt.Println(b)
	// a = append(a, b...) // แตก element แต่ละตัวของ b ใส่ต่อท้าย a (append(a, b[0], b[1]))
	a = append(a, b...)
	fmt.Println(a)
	a[6] = 100
	// แก้ไข element b ใน a, element ใน b ต้องเปลี่ยนด้วย x ก็จะเปลี่ยนด้วย เพราะมัน share memory กัน
	fmt.Printf("\n")
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	// คำถามคือ x มี cap แค่ 10 จะเกิดอะไรขึ้นถ้าเรา append a ไปอีก 4 element
	a = append(a, 101)
	a = append(a, 102)
	a = append(a, 103)
	a = append(a, 104)
	fmt.Printf("\n")
	fmt.Println(x) // x ไม่มีอะไรเปลี่ยนแปลง
	fmt.Println(a) // มีการจอง cap เพิ่มอีก 10 และทำการเพิ่ม element ใหม่
	fmt.Println(b)
	fmt.Println("len a is", len(a), ", cap a is", cap(a))
}
