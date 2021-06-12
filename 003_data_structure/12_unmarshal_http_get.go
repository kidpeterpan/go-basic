package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todos []struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
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
	// อ่าน stream response.Body
	respBd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("an error occurred when read resp body:", err)
		return
	}

	todos := Todos{}
	err = json.Unmarshal(respBd,&todos)
	if err != nil {
		fmt.Println("can not unmarshal todos:", err)
		return
	}
	fmt.Println(len(todos))
	fmt.Printf("todos: %+v", todos)
}
