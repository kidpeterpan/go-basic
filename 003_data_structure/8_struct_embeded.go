package main

import "fmt"

type People struct {
	name string
	age int
	Address // anonymouse field
}

type Address struct {
	zipcode string
}

func main(){
	// Embedded struct คือ struct ที่มี member เป็น struct
	// ใน go มี concept ของ anonymouse field
	addr := Address{ zipcode: "10260" }
	fmt.Printf("%+v\n",addr)
	pp := People{
		name: "Dang",
		age: 10,
	}	
	fmt.Printf("%+v", pp)
}
