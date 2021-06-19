package main

import "fmt"

type tmpStruct struct {
	name string
}

func (tmp *tmpStruct) SetNameByReference(newname string) {
	tmp.name = newname
}

func (tmp tmpStruct) SetNameByValue(newname string) {
	tmp.name = newname
}

func main() {
	tmp := tmpStruct{name: "Goku"}
	fmt.Println(tmp.name)
	tmp.SetNameByValue("Goku2")
	fmt.Println(tmp.name) // nothing change
	tmp.SetNameByReference("Goku3")
	fmt.Println(tmp.name)
}
