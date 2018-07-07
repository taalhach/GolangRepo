package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age int
}
type Student struct {
	Person
	RegNo string
	Enrollment int
}
type Employee struct {

	Person
	EmpID string
	Appointment int
}
func (p Person) BasicInformation() string{
	return "Name: "+p.Name+" Age: "+strconv.Itoa(p.Age)
}





func main()  {
	emp:=Employee{
		Person:Person{Name:"Ahmad ",Age:30},
		Appointment:2007,
		EmpID:"DNB-2007-109",
	}
	stdA:=Student{
		Person{Name:"talha",Age:25},
		"Sp15-BCS-042",
		2017,
	}
	// second way of initializing structs
	stdB:=Student{
		Person:Person{Name:"Ali Raza",Age:27},
		RegNo:"abc",
		Enrollment:2015,

	}
	fmt.Println(stdA.BasicInformation())
	fmt.Println(stdB.BasicInformation())
	fmt.Println("Employee info")
	fmt.Println(emp.BasicInformation())


}
