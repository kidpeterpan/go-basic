package main

import "fmt"

func main() {
	// zero value of slice is nil
	var slStr []string
	fmt.Println("slStr == nil:", slStr == nil) // true

	// literal slice -> get empty slice not nil slice
	slLiteral := []string{}
	fmt.Println(len(slLiteral)) // 0
	fmt.Println(cap(slLiteral))
	fmt.Println("slLiteral == nil:", slLiteral == nil) // false

	// literal slice with member
	sllmem := []string{"one", "two", "three"}
	fmt.Println(sllmem)     // [one two three]
	fmt.Println(sllmem[:2]) // should be [one two]
	fmt.Println(sllmem[1:]) // should be [two three]

	// ประกาศ slice ด้วย make()
	mSlice := make([]int, 10) // make([]int, 10, ใส่ cap ด้วยก็ได้)
	fmt.Println(mSlice)
	fmt.Println(len(mSlice)) // 10
	fmt.Println(cap(mSlice))
	fmt.Println("mSlice == nil:", mSlice == nil) // false
	// fmt.Println(mSlice[10]) // index out of range
	// การเข้าถึงด้วย sl[index] ยังไงก็ห้ามมากกว่า len แต่เราสามารถเพิ่ม member ได้ด้วย append()

	// ลองมาจอง mem เยอะๆ กัน
	// outOfMem := make([]int,999999999999)
	// fmt.Println(outOfMem)
}
