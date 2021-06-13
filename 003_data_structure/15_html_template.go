package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// note: ***html/template ช่วยเรา escape html script ให้เรา***

type TodoModel struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	// todos api https://jsonplaceholder.typicode.com/todos
	// call todos api and decode
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("an error occurred when call totos api:", err)
		return
	}
	jsonDecoder := json.NewDecoder(resp.Body)
	todos := []TodoModel{}
	err = jsonDecoder.Decode(&todos)
	if err != nil {
		fmt.Println("an error occurred when decode resp:", err)
		return
	}
	// สร้าง index.html
	// indexTmp = ioutil.ReadFile -> index.html
	idxTmp, err := ioutil.ReadFile("003_data_structure/index.html")
	if err != nil {
		fmt.Println("an error occurred when reading index.html:", err)
		return
	}
	tmp := template.Must(template.New("IndexTemplate").Parse(string(idxTmp)))
	_ = tmp.Execute(os.Stdout, todos)
	// run ด้วย go run 003_data_structure/15_html_template.go > 003_data_structure/output.html
	// note: ลองดู result ผ่าน browser
}
