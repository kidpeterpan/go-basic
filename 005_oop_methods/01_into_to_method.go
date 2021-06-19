package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) SetX(x int) {
	p.x = x
}

func main() {
	var p Point
	fmt.Println(p) // x 0, y 0
	p.SetX(10)
	fmt.Println(p.X()) // 10
}
