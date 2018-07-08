package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type message struct {
	UserName string `json:"userID"`
	Pass string `json:"password,omitempty"`
	Pin int `json:",omitempty"`
	Status string 	`json:"-"`   //skips this field
	Slashing string `json:"-,"`  // - is key
	Pincode float32 `json:",string"`   // marshals as string of json
}
var m =message{
	UserName:"taalhach",
	Pass:"",
	Pin:123,
	Status:"logged",
	Slashing:"dash value",
	Pincode:1.1,

}
var msg =message{
	UserName:"Escape",
	Pass:"pass123",
	Pin:321,
	Status:" not logged",
	Slashing:"new value",
	Pincode:1.1,

}
func main() {
	mdata,err:=json.Marshal(m)
	if err!=nil{
	log.Fatal("Marshalling failed")
	}
	fmt.Println(mdata)
	fmt.Println(string(mdata))

	mdata,err=json.MarshalIndent(msg,""," ")
	if err!=nil{
		log.Fatal("Marshalling failed")
	}
	fmt.Println(mdata)
	fmt.Println(string(mdata))





}
