package main

import (
	"encoding/json"
	"log"
	"fmt"
)

// in this example there is an anonymous field which contains same name of fields as outer struct contains


type Person struct {

	Name string `json:"user_name"`
	Age int `json:"Old"`
}
type student struct {
	Person
	Name string `json:"user_name"`
	Age int

}
type employee struct {
	Person
	EmpName string
	Old int
}

// least level nested fields are selected if same tags are assigned to outer and inner fields of strcuts
//https://golang.org/pkg/encoding/json/#Marshal


var emp =employee{
	Person:Person{Name:"Ali",Age:21},
	EmpName:"Dosa",
	Old:21,
}
var std =student{
	Person:Person{Name:"talha",Age:21},
	Name:"resilient",
	Age:21,
}

func main() {
	mdata,err :=json.MarshalIndent(std,""," ")
	if err!=nil{
		log.Fatal("Error in marshalling student")
	}
	fmt.Println(string(mdata))

	mdata,err =json.MarshalIndent(emp,""," ")
	if err!=nil{
		log.Fatal("Error in marshalling employee")
	}
	fmt.Println(string(mdata))
	}