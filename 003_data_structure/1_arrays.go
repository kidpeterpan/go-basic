package main

import (
	"fmt"
	"reflect"
)

// ***array is copy by value*** no share
func main() {
	// การเก็บข้อมูลแบบ'เรียงกัน'ลงใน ram
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits)) // 5

	// note - [2]int and [3]int are type different
	twoSlots := [2]int{}
	threeSlots := [3]int{}
	fmt.Println("Two slots:", reflect.TypeOf(twoSlots))
	fmt.Println("Three slots:", reflect.TypeOf(threeSlots))
	// fmt.Println("twoSlots == threeSlots:", twoSlots == threeSlots) //err miss match types

	copyFruits := fruits
	copyFruits[0] = "microsoft"
	fmt.Println(fruits[0])     // still be apple
	fmt.Println(copyFruits[0]) // microsoft
	// แล้วถ้าเราอยาก share ล่ะ
	pointerFruits := &fruits
	pointerFruits[0] = "amazon"
	fmt.Println(pointerFruits[0]) // should be amazon
	fmt.Println(fruits[0])        // should be amazon

	// note - ***ส่วนมากเราใช้ slice มากกว่า***

	// การประกาศ arrays แบบ ellisps
	cars := [...]string{2: "toyota", 0: "honda", 1: "benz"} // ใส่เพิ่มได้เรื่อยๆ กำหนด index ด้วย :
	fmt.Println(len(cars))
	fmt.Println(cars[0]) // should be "honda"
	cars2 := cars
	fmt.Println("cars2 == cars", cars2 == cars) // valid element & order should be true
	cars3 := cars
	cars3[0] = "toyota"
	cars3[2] = "honda"
	fmt.Println("cars == cars3", cars == cars3) // invalid order should be false

	// like other language go has multi dimensional
	arr2d := [2][2]int{}
	arr2d[0][0] = 1
	fmt.Println(arr2d)

}
