package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Orderqueue struct {
	Name string
	Orders map[int]string

}
// we can only use int and string as key of map in marshalling otherwise it throws an error
var m =make(map[int]string)

func main()  {
	m[1011]="order 1"
	m[101]="order 2"
	m[101]="order 3"
	m[101]="order 4"
	ent:=Orderqueue{
		Name:"AltEx",
		Orders:m,
	}
	mdata,err:=json.MarshalIndent(ent,""," ")
	if err!=nil{
		fmt.Println(err)
		log.Fatal("Error occured")

	}
	fmt.Println(string(mdata))

}
