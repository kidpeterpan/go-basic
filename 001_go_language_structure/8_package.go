package main

import (
	"fmt"
	"github.com/konghiran/001_go_language_structure/weight"
)

func main() {
	// สร้าง directory ชื่อเดียวกับ package ที่เราต้องการสร้าง
	// เช่น weight directory source code ด้านใน จะเป็น package weight
	// import  มาใช้ โดย เรียกผ่าน weight.CapitalFirstCharacterXXX <- export ใน go จะต้องขึ้นต้นด้วยตัวใหญ่
	kilo := weight.Kilo(10)
	pond := weight.KiloToPond(kilo)
	fmt.Printf("10 kilo is %f pond\n", pond)
}
