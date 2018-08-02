package main

import "fmt"

type Customer struct {
	Name,CusiD string
	Age int
}
type Emp struct {
	Name,EmpiD string
	Age int
}
func Accept(val interface{}) {
	fmt.Printf("%T",val)
	cus,err:=val.(Customer)
	if !err{
		fmt.Println("Not acceptable")
		return
	}
	fmt.Println(cus)
}
func main() {
	a:=Customer{Name:"Talha",CusiD:"a1",Age:12}
	Accept(a)
	b:=Emp{Name:"Talha",EmpiD:"a1",Age:12}
	Accept(b)

}
