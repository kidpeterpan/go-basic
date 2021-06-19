package main

import "fmt"

type Person struct {
	Name string
	Surname string
}

func (p Person) FullName() string {
	return p.Name + " " + p.Surname
}

type Employee struct {
	Id string
	Office string
	Person
}

func (e *Employee)Detail() string {
	return fmt.Sprintf("ID: %s, WORKING AT: %s, FULLNAME: %s",e.Id,e.Office,e.FullName())
}
func (e *Employee) IsSameOffice(anotherEmployee *Employee) bool {
	return e.Office == anotherEmployee.Office
}

type Programmer struct {
	Language []string
	Employee
}

func(p *Programmer)Detail()string {
	return fmt.Sprintf("ID: %s, WORKING AT: %s, FULLNAME: %s, USE LANGUAGE: %+v",p.Id,p.Office,p.FullName(),p.Language)
}

type Tester struct {
	Tools []string
	Employee
}

func (t Tester) Detail() string {
	return fmt.Sprintf("ID: %s, WORKING AT: %s, FULLNAME: %s, USE TOOL: %+v",t.Id,t.Office,t.FullName(),t.Tools)
}
func main() {
	peepee := Employee{
		Id: "A123",
		Office: "Thailand",
		Person: Person{
			Name: "Peepee",
			Surname: "David",
		},
	}
	feefee:= Employee{
		Id: "A124",
		Office: "Thailand",
		Person: Person{
			Name: "Feefee",
			Surname: "John",
		},
	}
	fmt.Println(peepee.IsSameOffice(&feefee))
	peepeeAsProg := Programmer{
		Language: []string{"Go","Kotlin","Typescript"},
		Employee: peepee,
	}
	feefeeAsTester := Tester{
		Tools: []string{"Robot"},
		Employee: feefee,
	}
	fmt.Println(peepeeAsProg.Detail())
	fmt.Println(feefeeAsTester.Detail())
	fmt.Println(peepeeAsProg.IsSameOffice(&feefeeAsTester.Employee))
}
