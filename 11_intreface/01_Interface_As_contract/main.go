package main

import (
	"strconv"
	"fmt"
)
type Contract interface {
	String() string
}

type Customer struct {
	Name,CusiD string
	Age int
}

func (c Customer) String()string  {
	return c.Name+c.CusiD+strconv.Itoa(c.Age)
}

type Emp struct {
	Name,EmpiD string
	Age int
}

func (e Emp) String()string  {
	return e.Name+e.EmpiD+strconv.Itoa(e.Age)
}

func Details(c Contract)  {
	fmt.Println(c.String())
}

func main() {
	a:=Customer{Name:"Talha ",CusiD:"a1 ",Age:12}
	b:=Emp{Name:"Asad ",EmpiD:"b1 ",Age:12}
	Details(a)
	Details(b)
}