package main

import (
	"fmt"
)

// note: variable and func member can not use the same name
type animal struct {
	Name  string
	Sound string
}

//DoSound is method, method is behavior of objects
func (a *animal) DoSound() { // (a *animal) this called reciever"
	fmt.Println("Sound:", a.Sound)
}

func (a *animal) ChangeName(newname string) {
	a.Name = newname
}

func main() {
	cat := animal{
		Name:  "Mubin",
		Sound: "Meow",
	}
	cat.DoSound()
	fmt.Println("his name is:", cat.Name)
	cat.ChangeName("Mudum")
	fmt.Println("his name is:", cat.Name)
}
