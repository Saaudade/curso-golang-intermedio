package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporatyTimeEmployee struct {
	Person
	Employee
	taxRate int
}

func (tpEmployee TemporatyTimeEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 22
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	// GetMessage(ftEmployee) Esto no es permitido porque no es una Herencia si no una composición
	tEmployee := TemporatyTimeEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
