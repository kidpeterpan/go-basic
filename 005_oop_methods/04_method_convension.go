package main

import "fmt"

// note: ถ้า receiver เป็น pointer ก็ให้เป็น pointer ให้หมด

type Convension struct {
	rule string
}

func (c *Convension) SetRule(rule string) {
	c.rule = rule
}

func (c *Convension) Rule() string {
	return c.rule
}

func main() {
	var conv Convension
	conv.SetRule("If apply pointer to receivers should apply to all of them")
	fmt.Println(conv.Rule())
}
