package main

import (
	"fmt"
	"sort"
)

type PersonForSort []struct {
	name string
	age  int
}

//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}

func (pfs PersonForSort) Len() int {
	return len(pfs)
}

func (pfs PersonForSort) Less(i, j int) bool {
	return pfs[i].name < pfs[j].name
}

func (pfs PersonForSort) Swap(i, j int) {
	pfs[i], pfs[j] = pfs[j], pfs[i]
}

func main() {
	// sort Person by name
	// input is []Person
	pfs := PersonForSort{
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
	sort.Sort(&pfs)
	fmt.Println(pfs)
}
