package main

import (
	"encoding/json"
	"fmt"
)

type InvalidTodo struct {
	userId    int
	id        int
	title     string
	completed bool
}

type Todo struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	// we will generate json from https://jsonplaceholder.typicode.com/

	// Marshal คือการจัดระเบียบ จาก memory 010101 -> File/string/stream
	// Unmarshal คือทำย้อนกลับ จาก File/string/stream -> memory 0101001
	data := `
	{
		"userId": 1,
		"id": 2,
		"title": "quis ut nam facilis et officia qui",
		"completed": false
	}
	`
	invalidTodo := InvalidTodo{}
	err := json.Unmarshal([]byte(data), &invalidTodo)
	if err != nil {
		fmt.Println("can not unmarshal", err)
		return
	}
	fmt.Printf("\n %+v", invalidTodo) // ได้ zero value เราทำอะไรผิด??? -> เราไม่ได้ export field

	todo := Todo{}
	err = json.Unmarshal([]byte(data), &todo)
	if err != nil {
		fmt.Println("can not unmarshal", err)
		return
	}
	fmt.Printf("\n %+v", todo) // unmarshal correct with export field
	// {UserId:1 Id:2 Title:quis ut nam facilis et officia qui Completed:false}
}
