package main

import "fmt"

type Human struct {
	name string
	age  int
}

// note: **struct copy by value**
func main() {
	// ถ้าเราไม่ assign ค่าให้กับ member ของ struct จะได้ zero value ของ type นั้นๆ
	hu := Human{}
	fmt.Println(hu) // { name: "", age: 0}

	hu1 := Human{name: "deku", age: 14}
	hu2 := hu1 // **copy by value**
	hu2.name = "uraraka"
	fmt.Printf("\n")
	fmt.Println(hu1) // { name: "deku", age: 14}
	fmt.Println(hu2) // { name: "uraraka", age: 14}

	// ถ้าเราอยากให้ hu3 reference ไปที่เดียวกับ hu2 เราใช้ pointer ช่วย
	hu3 := &hu2  // copy address of hu2 to h3
	hu3.age = 13 // (*h3).age = 13 -> hu2.age ควรเปลี่ยนด้วย
	// บรรทัดข้างบน compiler ช่วยเราในการ implicit dereference
	fmt.Printf("\n")
	fmt.Println(hu2)  // { name: "uraraka", age: 13 }
	fmt.Println(*hu3) // { name: "uraraka", age: 13 }

	// compare struct
	fmt.Printf("\n")
	fmt.Println("hu1 == hu2?", hu1 == hu2)
	fmt.Println("hu2 == hu3?", hu2 == *hu3) // dereference hu3 from pointer to struct
}
