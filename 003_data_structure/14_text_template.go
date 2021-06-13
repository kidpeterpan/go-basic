package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Item struct {
	Name  string // ต้อง Export ด้วย เพราะถูกเรียกใช้จาก text/template pkg
	Count int
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {
	items := []Item{
		{Name: "bottle", Count: 2},
		{Name: "mouse", Count: 3},
	}
	const itemDesTmp = `the item name {{ .Name }}, have {{ .Count }} {{ "\n" }}` // . is "Item", ส่วน \n มองเป็น string เลย require ""
	// {{ }} สิ่งที่อยู่ด้านในเรียกว่า actions
	tmp, err := template.New("ItemDescription").Parse(itemDesTmp)
	// ใช้ template.Must ได้ถ้าขี้เกียจ handle err
	if err != nil {
		fmt.Println("an error occurred when parse the template:", err)
		return
	}
	fmt.Println(tmp.Name())
	_ = tmp.Execute(os.Stdout, items[0])
	_ = tmp.Execute(os.Stdout, items[1])

	// template 'RANGE' with arrays
	fmt.Printf("\n")
	fmt.Println("==== template 'RANGE' with arrays")
	const itemDesTmpRange = `{{ range . }}the item name {{ .Name | upperCase }}, have {{ .Count }} {{ "\n" }}{{ end }}`
	// note: ใช้ | เพื่อเรียก func ที่ register แล้ว
	fm := make(template.FuncMap)
	fm["upperCase"] = upperCase
	tmpRange := template.Must(template.New("ItemDescriptionWithRange").
		Funcs(fm). // register upperCase func
		Parse(itemDesTmpRange))
	fmt.Println(tmp.Name())
	_ = tmpRange.Execute(os.Stdout, items)
}
