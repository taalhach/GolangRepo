package main

import "fmt"

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





func main()  {
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

	fmt.Println(stdA.Name,stdA.Age)

	fmt.Println(stdA)
	fmt.Println(stdB)

}
