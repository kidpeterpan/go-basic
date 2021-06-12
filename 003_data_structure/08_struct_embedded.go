package main

import "fmt"

type People struct {
	// note: หากต้องการ export field ให้ตั้งชื่อขึ้นต้นด้วยพิมพ์ใหญ่
	name    string
	age     int
	Address // anonymous field
}

type Address struct {
	zipcode string
}

func main() {
	// Embedded struct คือ struct ที่มี member เป็น struct
	// ใน go มี concept ของ anonymous field
	addr := Address{zipcode: "10260"}
	fmt.Printf("%+v\n", addr)
	pp := People{ // because of we can not mix แบบระบุ field กับไม่ระบุ field ใน literal ได้
		name: "Dang",
		age:  10,
	}
	pp.zipcode = addr.zipcode // สามารถเรียกใช้ pp.zipcode ได้เลย ไม่ต้อง pp.Address.zipcode
	fmt.Printf("%+v", pp)

	// ถ้าจะประกาศให้ assign zipcode ใน People ได้เลย
	pp2 := People{
		"Dang",
		10,
		Address{zipcode: "10260"},
	}
	fmt.Printf("\n%+v", pp2)
	fmt.Printf("\n%s", pp2.zipcode)         // 10260
	fmt.Printf("\n%s", pp2.Address.zipcode) // หรือเราจะเรียก pp2.Address.zipcode ก็ทำได้ แต่มันยาว
}
