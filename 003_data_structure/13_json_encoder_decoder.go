package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Todos2 []struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// json.Encode -> เอาอะไรที่ไม่รู้จัก encode ไปในรูปของ json
	// json.Decode -> เรารู้จักว่า json เป็นยังไง แต่คนอื่นไม่รู้จัก เราจะต้อง decode เป็น struct
	// ex. file reader -> json decoder -> struct instance
	// call todos api
	// https://jsonplaceholder.typicode.com/todos
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	// close body
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("an error occurred when call todos api:", err)
		return
	}
	// Decode response.Body
	fmt.Println("=== Decode")
	jsonDecoder := json.NewDecoder(resp.Body)
	todos := Todos2{}
	if err := jsonDecoder.Decode(&todos); err != nil {
		fmt.Println("an error occurred when decode resp:", err)
		return
	}
	fmt.Println(len(todos))
	fmt.Printf("todos: %+v\n", todos)

	// Encode to Json
	// note: file ในความหมายของ unix รวมถึงหน้าจอด้วย
	fmt.Println("=== Encode")
	jsonEncoder := json.NewEncoder(os.Stdout)
	if err := jsonEncoder.Encode(todos); err != nil {
		fmt.Println("an error occurred when encode todo to json:", err)
		return
	}
}
