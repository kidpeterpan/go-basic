package main

import "fmt"

// string เป็น immutable (เปลี่ยนแปลงไม่ได้)
func main() {
	eng := "hello"
	th := "สวัสดี"
	fmt.Println(len(eng)) // 5
	fmt.Println(len(th))  // 18 ทำไมได้ 18 หว่า??? เพราะว่า len ของ string คือ len ของ string in byte

	fmt.Println(string([]byte{65})) // A จริงๆ แล้ว string ก็คือ []byte{}

	// hi สวัสดี
	testByte := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	fmt.Println(string(testByte[0])) // h
	fmt.Println(string(testByte[1])) // i
	fmt.Println(string(testByte[2])) // -
	fmt.Println(string(testByte[3])) // à กลับไม่ใช่ ส แบบที่คาดหวัง

	// unicode -> enable people around the world to use computers in any language https://unicodelookup.com/
	fmt.Println(string(rune(0xE2A))) // ส -> จะสังเกตุว่า testByte ของเราไม่มี 0xE2A เลย แล้วมันปริ้น ส ได้ไง

	// แปลง unicode to UTF-8
	// E2A ที่เป็น hex ถ้าเป็น binary จะได้ 111000101010
	// จับเข้าสูตร range 12 ของ UTF-8 จะได้ 1110xxxx 10xxxxxx 10xxxxxx -> 11100000 10111000 10101010
	// เสร็จแล้วแปลงกลับเป็น hex จะได้ E0 B8 AA

	// ใน go ใช้ UTF-8
	// จริงๆ แล้ว ส ก็คือ
	fmt.Println("UTF-8:", string(testByte[3:6])) // e0, b8, aa -> ส

	// ถ้าอยากได้ len ที่ถูกต้องของภาษาไทย แปลง string -> []rune{} ก่อน
	fmt.Println(len([]rune(th)))
}
