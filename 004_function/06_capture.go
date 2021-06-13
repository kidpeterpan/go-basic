package main

import "fmt"

// createListFunc return function that can print single value of numbers input
// คือ เรียก แต่ละครั้งก็จะ print member ของ input list
func createListFunc(numbers []int) []func() {
	listFunc := []func(){}
	for _, value := range numbers {
		capture := value
		listFunc = append(listFunc, func() {
			// value ในนี้ให้เรามองเป็น global value
			// ดังนั้นถ้าเราจะแก้ปัญหาได้ 0 หมด ต้องหา variable closure อีกตัวมา capture มัน
			// fmt.Println(value)
			// ที่ถูกต้องคือ
			fmt.Println(capture)
		})
	}
	return listFunc
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fns := createListFunc(numbers)
	for _, fn := range fns {
		fn()
	}
}
