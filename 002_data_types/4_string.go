package main

import "fmt"

// string เป็น immutable (เปลี่ยนแปลงไม่ได้)
func main() {
	eng := "hello"
	th := "สวัสดี"
	fmt.Println(len(eng)) // 5
	fmt.Println(len(th)) // 18 ทำไมได้ 18 หว่า??? เพราะว่า len ของ string คือ len ของ string in byte

	fmt.Println(string([]byte{65})) // A จริงๆ แล้ว string ก็คือ []byte{}

	// hi สวัสดี
	testByte := []byte{0x68,0x69,0x2d,0xe0,0xb8,0xaa,0xe0,0xb8,0xa7,0xe0,0xb8,0xb1,0xe0,0xb8,0xaa,0xe0,0xb8,0x94,0xe0,0xb8,0xb5}
	fmt.Println(string(testByte[0])) // h
	fmt.Println(string(testByte[1])) // i
	fmt.Println(string(testByte[2])) // -
	fmt.Println(string(testByte[3])) // à กลับไม่ใช่ ส แบบที่คาดหวัง

}
