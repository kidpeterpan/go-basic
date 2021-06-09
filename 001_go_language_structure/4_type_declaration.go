package main

import "fmt"

type Kilo float64
type Pond float64

func (p Pond) toKilo() Kilo {
	k := p / 2.2046226218 // go implicit 2.20... to Pond(2.20...)
	fmt.Printf("k's type is: %T\n", k)
	return Kilo(k) // then we cast pond to Kilo so we can compare Kilo vs Kilo
}

func main() {
	// เช่น เรามี weight = 5 จะไม่รู้ว่ามันเป็นกิโลหรือปอน
	weight := 5
	fmt.Println("what is the unit of weight:", weight, "kg or pond?")
	// เหมือนว่าทีมผู้สร้าง Go จะเล็งเห็นปัญหานี้เลย อนุญาติให้เราสร้าง type เองได้

	// fmt.Println(Kilo(10) == Pond(10)) เราไม่สามารถนำ type Kilo มาเทียบกับ Pond ได้เพราะถือว่าเป็นคนละ type กัน
	// แต่ลองมาดู
	fmt.Println("is 10 Kilo == 10?:", Kilo(10) == 10)
	// สาเหตุที่มัน compare กับ 10 ได้เพราะว่าเราไม่ได้ประกาศชัดเจนว่า 10 เป็น type อะไร go จะทำสิ่งที่เรียกว่า
	// implicit declaration ดูว่า 10 ของเราทำเป็น Kilo(10) ได้ไหม ถ้าไม่ได้จะ compilation error
	// ดังนั้น 10 ในที่นี้ go มองเป็น Kilo(10) นี่แหละ

	// ในเมื่อคนละ type compare กันไม่ได้ต่อให้ cast ค่าก็ผิด, ดังนั้นเราต้องการ method เพื่อแปลงจาก Pond to Kilo
	fmt.Println("2 Kilo > 2 Pond : ", Kilo(2) > Pond(2).toKilo())
	fmt.Println("because 1 pond is:", Pond(1).toKilo(), "Kilo")
}
