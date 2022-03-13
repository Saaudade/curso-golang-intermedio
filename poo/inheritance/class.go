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
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 22
	ftEmployee.id = 1
	fmt.Printf("%v", ftEmployee)
	// GetMessage(ftEmployee) Esto no es permitido porque no es una Herencia si no una composición
}
