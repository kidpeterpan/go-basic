package main

import "fmt"

type NilReceiver struct {
	name string
}

func (nr *NilReceiver) SayHi() {
	fmt.Println("Hi! Printed from NilReceiver")
}

func (nr *NilReceiver) SayHello() {
	fmt.Println("Hello! Printed from NilReceiver", nr.name)
}

func main() {
	var nr *NilReceiver
	nr = nil
	fmt.Println(nr)
	nr.SayHi()
	nr.SayHello() // error เพราะว่าเรามีการเรียกใช้ nil.name
}
