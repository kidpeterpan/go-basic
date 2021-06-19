package main

import (
	"flag"
	"fmt"
	"strconv"
)

// ถ้ารันโดยไม่ใส่ flag -name=xxx จะคืน anonymous เป็น default
var name = flag.String("name", "anonymous", "Your name")
var flagRomanAge = func() *RomanAge {
	ra := RomanAge{}
	flag.Var(&ra, "age", "Your age in roman")
	return &ra
}
var age = flagRomanAge()

// โจทย์คือเราอยากได้ flag -age=III แล้วจะได้ value 3
// type Value interface {
//    String() string
//    Set(string) error
// }
// เราจะ implement Value แล้วใช้กับ flag.Var() เพื่อสร้าง custom flag

type RomanAge struct {
	age string
}

func (r *RomanAge) String() string {
	var mapRomanAge = make(map[string]int)
	mapRomanAge["i"] = 1
	mapRomanAge["ii"] = 2
	mapRomanAge["iii"] = 3
	mapRomanAge["iv"] = 4
	mapRomanAge["v"] = 5
	return strconv.Itoa(mapRomanAge[r.age])
}

func (r *RomanAge) Set(str string) error {
	r.age = str
	return nil
}

func main() {
	// flag.Parse() ช่วยเราในการ parse arguments
	flag.Parse()
	fmt.Println(*name)
	// run ด้วย go run 006_oop_interface/03_flag_interface.go -name=Peepee
	fmt.Println(age) // 3
}
