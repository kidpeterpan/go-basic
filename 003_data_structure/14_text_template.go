package main

import (
	"fmt"
	"os"
	"text/template"
)

type Item struct {
	Name  string // ต้อง Export ด้วย เพราะถูกเรียกใช้จาก text/template pkg
	Count int
}

func main() {
	items := []Item{
		{Name: "bottle", Count: 2},
		{Name: "mouse", Count: 3},
	}
	const itemDesTmp = `the item name {{ .Name }}, have {{ .Count }} {{ "\n" }}` // . is "Item", ส่วน \n มองเป็น string เลย require ""
	tmp, err := template.New("ItemDescription").Parse(itemDesTmp)
	// ใช้ template.Must ได้ถ้าขี้เกียจ handle err
	if err != nil {
		fmt.Println("an error occurred when parse the template:", err)
		return
	}
	fmt.Println(tmp.Name())
	_ = tmp.Execute(os.Stdout, items[0])
	_ = tmp.Execute(os.Stdout, items[1])

}
