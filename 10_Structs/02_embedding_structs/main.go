package main

import "fmt"

type Person struct {
	Name string
	age int32
}
type Employee struct {
	P Person
	emp_id int32
	duty_hrs string
}

func main() {
	emp:=Employee{
		P:Person{Name:"talha",age:56},
		emp_id:123,
		duty_hrs:"8 to 5",
	}
	fmt.Println(emp)
}

