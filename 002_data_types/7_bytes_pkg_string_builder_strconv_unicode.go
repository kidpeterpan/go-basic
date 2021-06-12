package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	x := "ทดสอบทดสอบทดสอบทดสอบทดสอบทดสอบ"
	finder := "สอ"
	found := bytes.Count([]byte(x), []byte(finder))
	fmt.Println(found) // 6
	// ใช้ strings package ก็ได้ไม่ต้อง cast
	found = strings.Count(x, finder)
	fmt.Println(found) // 6

	buff := bytes.Buffer{}
	buff.WriteString("Y")
	buff.WriteString("o")
	fmt.Println(buff.String())

	// ถ้าจะใช้ bytes.Buffer ใช้ strings.Builder ดีกว่า
	strBd := strings.Builder{}
	strBd.WriteString("Y")
	strBd.WriteString("o")
	fmt.Println(strBd.String())

	// แปลง string เป็น integer
	result, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("err: %s \n", err)
		return
	}
	fmt.Println(result, reflect.TypeOf(result)) // 123 int

	result2 := strconv.Itoa(456)
	fmt.Println(result2, reflect.TypeOf(result2)) // "456" string

	// strconv.Parse...
	fmt.Println(strconv.ParseBool("True"))  // true
	fmt.Println(strconv.ParseBool("true"))  // true
	fmt.Println(strconv.ParseBool("x"))     // err
	fmt.Println(strconv.ParseBool("False")) // false
	fmt.Println(strconv.ParseBool("false")) // false
	fmt.Println(strconv.ParseBool(""))      // err
	fmt.Println(strconv.ParseBool("0"))     // false
	fmt.Println(strconv.ParseBool("1"))     // true

	// play with unicode package
	fmt.Println(unicode.IsUpper('A')) // true

}
