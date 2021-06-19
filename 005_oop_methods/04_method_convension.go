package main

import "fmt"

// note: ถ้า receiver เป็น pointer ก็ให้เป็น pointer ให้หมด

type Convention struct {
	rule string
}

func (c *Convention) SetRule(rule string) {
	c.rule = rule
}

func (c *Convention) Rule() string {
	return c.rule
}

func main() {
	var conv Convention
	conv.SetRule("If apply pointer to receivers should apply to all of them")
	fmt.Println(conv.Rule())
}
