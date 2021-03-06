package main

import "fmt"

func main() {
	x := 2
	var interfaceX interface{}
	interfaceX = 2
	switchWithValue(x)
	switchWithCondition(x)
	switchWithType(interfaceX)
}

func switchWithType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Println("x is unknown")
	}
}

func switchWithCondition(x int) {
	switch { // จริงๆ ตรงนี้คือ switch true { } แต่ว่าเรา omit มัน
	case x == 2:
		fmt.Println("x is 2")
	case x == 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is unknown")
	}
}

func switchWithValue(x int) {
	switch x {
	case 2:
		fmt.Println("x is 2")
		// ใน go จะ break เลย ถ้าอยากไหลไป case ถัดไปให้ใส่ fallthrough
	case 3:
		fmt.Println("x is 3")
	default:
		// เหมือนภาษาอื่นๆ ถ้าไม่เข้า case ไหนเลย จะตก default
		fmt.Println("x is unknown")
	}
}
