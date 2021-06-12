package main

import (
	"fmt"
)

// maps -> reference to hash table
// key จะเป็นอะไรก็ได้ที่สามารถนำมา compare ได้ ( a == b )
func main() {
	// ประกาศ map แบบ literal
	mapStrInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(mapStrInt)
	fmt.Printf("\n")
	// ประกาศ map ด้วย make()
	mapStrIntWithMake := make(map[string]int)
	mapStrIntWithMake["one"] = 1
	mapStrIntWithMake["two"] = 2
	mapStrIntWithMake["three"] = 3
	fmt.Println(mapStrIntWithMake)
	fmt.Printf("\n")
	// การเข้าถึง map element
	fmt.Println("mapStrInt[\"one\"] is", mapStrInt["one"])
	fmt.Printf("\n")
	// test map is reference type
	refMapStrInt := mapStrInt
	// ถ้ามีการเพิ่ม key ที่ refMapStrInt ที่ mapStrInt ควรเพิ่มด้วย
	refMapStrInt["four"] = 4
	fmt.Println(refMapStrInt)
	fmt.Println(mapStrInt)
	fmt.Printf("\n")
	// delete element from map
	delete(refMapStrInt, "four")
	fmt.Println(refMapStrInt)
	fmt.Println(mapStrInt)
	fmt.Printf("\n")
	// check ว่ามี key นี้ใน map ไหม
	// map จะคืนค่าสองค่า คือ value กับ ok
	value, ok := mapStrInt["four"]
	if !ok {
		fmt.Println("not found key 'four'")
	} else {
		fmt.Println("value of mapStrInt[\"four\"] is", value)
	}
	// note: **map not garantee order**
	fmt.Printf("\n")
	fmt.Println("len of mapStrInt is", len(mapStrInt))
	for key, value := range mapStrInt {
		// run แต่ละครั้ง order อาจจะแตกต่างกัน
		fmt.Printf("mapStrInt[\"%s\"] is %d\n", key, value)
	}
	// ถ้าเราพยายามเข้าถึง key ที่ไม่มีอยู่จะไม่ error แต่ได้ zero value แทน
	fmt.Printf("\n")
	fmt.Println(mapStrInt["four"])
	fmt.Println(mapStrInt["five"])
}
