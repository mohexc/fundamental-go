package gobasic

import "fmt"

// create a struct
type person struct {
	name string
	age  int
}
type employee struct {
	person
	emplID int
}

// method of struct
func (p person) detail() {
	fmt.Println("---- detail of person: ----")
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
}

func (e employee) detail() {
	fmt.Println("---- detail of employee: ----")
	fmt.Println("name:", e.name)
	fmt.Println("age:", e.age)
	fmt.Println("emplID:", e.emplID)
}

func ShowStruct() {
	// create a instance of person
	p := person{name: "John", age: 30}
	fmt.Println("create instance of person:", p)

	// delete instance of person
	p = person{}
	fmt.Println("delete instance of person:", p)

	// update instance of person
	p.name = "John"
	p.age = 30
	p.detail()

	// create a instance of employee
	e := employee{person: person{name: "John", age: 30}, emplID: 1}
	e.detail()

}
