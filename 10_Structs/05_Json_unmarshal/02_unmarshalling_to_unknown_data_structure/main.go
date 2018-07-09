package main

import (
	"encoding/json"
	"log"
	"fmt"
)

var b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
var ds interface{}

func main() {
	err :=json.Unmarshal(b,&ds)
	if err!=nil{
		log.Fatal("Error in unmarshaling",err)
	}
	fmt.Printf("%T\n",ds)
	fmt.Println(ds)

}
