package main

import "fmt"

func main() {
	// ใน go มีแค่ for loop
	for i:=0; i < 10; i++ {

		if i == 3 {
			fmt.Println("we no need 3")
			continue
		}

		if i == 8 {
			fmt.Println("the program need terminated at 8")
			break
		}

		fmt.Println(i)
	}

	// เราสามารถมีแค่ส่วนที่เป็น condition ได้ เช่น
	//for true {
	//	// infinite loop ...
	//}
}
