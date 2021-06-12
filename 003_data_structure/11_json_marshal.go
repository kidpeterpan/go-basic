package main

import (
	"encoding/json"
	"fmt"
)

type Todo2 struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"` // เราต้องการให้ show true/false ได้อย่างถูกต้อง
}

func main() {
	data := `
	[{
		"userId": 1,
		"id": 2,
		"title": "Hi there!",
		"completed": false
	},{
		"userId": 2,
		"id": 3,
		"completed": false
	},{
		"userId": 3,
		"id": 4,
		"title": "Yoh!" 
	}]
	`
	todos := []Todo2{}
	err := json.Unmarshal([]byte(data), &todos)
	if err != nil {
		fmt.Println("can not unmarshal", err)
		return
	}

	completed := true
	todos[0].Completed = &completed

	// ต่อไปเราจะทำการ marshal กลับเป็น json
	result, err := json.MarshalIndent(todos, "", "    ") // จริงๆ ใช้ Marshal ก็ได้ แต่เพื่อ format ที่สวยงาม
	if err != nil {
		return
	}
	fmt.Println(string(result)) // string([]byte) -> ได้ json กลับมาแล้ว

	// แต่ field name ขึ้นต้นด้วยตัวใหญ่หมดเลย !!!
	// go มี  struct tag เช่น User Id int `json:"userId"` มาจัดกันเลย
	// แค่นี้ result ก็จะถูกต้องแบบที่ต้องการแล้ว
}
