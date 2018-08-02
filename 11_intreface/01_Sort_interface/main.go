package main

import (
	"sort"
	"fmt"
)

type Customer struct {
	name, empId string
	age  int
}

type CustomerData []Customer

func (C CustomerData) Len() int {
	return len(C)
}
func (C CustomerData)	Less(i,j int) bool {
	return C[i].name<C[j].name
}
func (C CustomerData)	Swap(i,j int)  {
	 C[i],C[j]=C[j],C[i]
}

func main() {
	c:=CustomerData{
		{
			name:"Talha",
			empId:"a1",
			age:23,
		},{
			name:"Asad",
			empId:"b1",
			age:16,
		},
	}
	sort.Sort(c)
	fmt.Println(c)
}