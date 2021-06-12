package main

import (
	"fmt"
	"reflect"
)

func main() {
	// const ประกาศแล้วจะ reassign ไม่ได้ระหว่าง execution
	//var x = 2
	//var y = 0
	//fmt.Println(x / y) // runtime err (divide by zero)
	//fmt.Println(x / 0) // compile err (divide by zero)

	// สำหรับ const, evaluation เกิดตอน compile time
	// ถ้า debug จะต้องกระโดดข้าม line ของ const เพราะมันทำไปแล้วตอน compile time

	// untype const
	const persons = 4
	toffee := 5 / persons // implicite to int
	cost := 2.0 / persons // implicite to float64

	fmt.Println(toffee, reflect.TypeOf(toffee))
	fmt.Println(cost, reflect.TypeOf(cost))

	// group const
	const (
		Mon = iota // start with 0
		Tu
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	// ถ้าเราใส่แค่ Mon = 1 มันจะ copy ค่าทุกตัวจะเท่ากับ 1
	// ดังนั้น iota จะมาช่วยเรา, iota คือ constant generator
	fmt.Println(Mon, Tu, Wed, Thu, Fri, Sat, Sun) // 0 1 2 3 4 5 6
}
