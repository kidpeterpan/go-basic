package main

import "fmt"

type PersonForSort struct {
	name string
	age  int
}

func main() {
	// sort Person by name
	// input is []*Person
	pfs := []*PersonForSort{
		{name: "A", age: 11},
		{name: "A", age: 20},
		{name: "A", age: 10},
		{name: "B", age: 10},
		{name: "B", age: 9},
		{name: "D", age: 10},
		{name: "E", age: 10},
		{name: "C", age: 10},
		{name: "G", age: 10},
		{name: "F", age: 10},
	}
	fmt.Println("no sort ===")
	for _, pf := range pfs {
		fmt.Printf("%s, %d\n", pf.name, pf.age)
	}
	// implement interface sort.Interface and use sort.Sort
	fmt.Println("sorted ===")
}
