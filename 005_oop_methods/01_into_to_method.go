package main

import "fmt"

type point struct {
	x int
	y int
}

func (p *point) X() int {
	return p.x
}

func (p *point) SetX(x int) {
	p.x = x
}

func main() {
	var p point
	fmt.Println(p) // x 0, y 0
	p.SetX(10)
	fmt.Println(p.X()) // 10
}
