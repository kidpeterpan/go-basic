package main

import (
	"fmt"
	"math"
)

func main() {
	// int ใน go จะแบ่งเป็น sign ค่า + -
	// go use two's complement ในการ represents negative integer so int8's range is between -128 to 127
	fmt.Println("range of int8:",math.MinInt8, math.MaxInt8)
	fmt.Println("range of int16:",math.MinInt16, math.MaxInt16)
	fmt.Println("range of int32:",math.MinInt32, math.MaxInt32)
	var x int = 10
	fmt.Println("x is:",x)
	// unsigned ค่า + เท่านั้น
	fmt.Println("range of uint8: 0",math.MaxUint8)
	fmt.Println("range of uint16: 0",math.MaxUint16)
	fmt.Println("range of uint32: 0",math.MaxUint32)
	//var y uint = -1
	//fmt.Println(y) constant -1 overflows uint
}
