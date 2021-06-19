package main

import "fmt"

// note: ใน go ถ้า struct มี func signature ครบตามที่ interface มีถือว่า implement interface
func main() {
	// sword, bow, implement weapon
	var w weapon
	// call attack!
	// interface ก็มีการ embedded ได้เหมือนกัน (add interface Item { cost() }
	w = &sword{ name: "ice sword" }
	printAttack(w)
	w = &bow{ name: "fire bow" }
	printAttack(w)
}

func printAttack(w weapon) {
	fmt.Printf("attack: %s, cost: %d\n",w.attack(), w.cost())
}

type item interface {
	cost() int
}

type weapon interface {
	attack() string
	item
}

type sword struct {
	name string
}

func (s *sword) attack() string {
	return fmt.Sprintf("attack with: %s",s.name)
}

func (s *sword) cost() int {
	return 10
}

type bow struct {
	name string
}

func (b *bow) attack() string {
	return fmt.Sprintf("attack with: %s",b.name)
}

func (b *bow) cost() int {
	return 12
}


