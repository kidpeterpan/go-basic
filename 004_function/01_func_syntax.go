package main

import (
	"fmt"
	"reflect"
)

func avg(a float64, b float64) float64 { // ใน go return value มีได้มากกว่า 1
	return (a + b) / 2
}

// note: *** ใน go ไม่มี concept ของ overload และ default param ***
//func avg(a float64, b float64, c float64) float64 {
//	return ( a + b )/ 2
//}

func main() {
	// a := (1 + 3) / 2
	// b := (2 + 4) / 2
	// ถ้าเราใช้ func จะ save กว่า ความผิดพลาดน้อยกว่า แก้ไขที่เดียว
	a := avg(1, 3)
	b := avg(2, 4)
	fmt.Printf("a is %.2f, b is %.2f \n", a, b) // a is 2.00, b is 3.00
	fmt.Printf("func signature is -> %s \n", reflect.TypeOf(avg))
}
