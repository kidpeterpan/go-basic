package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// member ของ struct มี type ที่แตกต่างกันได้เช่น
	deku := Person{name: "deku", age: 14}
	fmt.Println(deku.name)
	fmt.Println(deku.age)
	fmt.Printf("\n")

	// ประกาศ struct แบบ literal
	uraraka := struct {
		name string
		age  int
	}{
		name: "uraraka",
		age:  14,
	}
	fmt.Println(uraraka.name)
	fmt.Println(uraraka.age)
	fmt.Printf("\n")
}
