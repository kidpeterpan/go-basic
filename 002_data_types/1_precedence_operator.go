package main

import "fmt"

func main() {
	// 5 * / % << >> & &^
	// 4 + - | ^
	// 3 == != < <= > >=
	// 2 &&
	// 1 ||
	// แต่เวลาใช้จริงใส่ ( ) ให้หน่อยดีกว่า คนอ่านจะได้ไม่งง
	x := 3*2 + 1
	fmt.Println("x should be 7 not 9:", x == 7)
}
