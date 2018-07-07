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

	fmt.Println(stdA.Name,stdA.Age)

	fmt.Println(stdA)

}
