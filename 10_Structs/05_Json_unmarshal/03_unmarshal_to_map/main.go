package main

import (
	"log"
	"encoding/json"
	"fmt"
)

var m =make(map[string]int)
func marshalMap() []byte {
	m["taalha"]=0123
	m["ali"]=0147
	mdata,err:=json.Marshal(m)
	if err!=nil{
		log.Fatal("Error in marshaling: ",err)
	}
	return mdata

}

var umdata interface{
}

func main() {

	err:=json.Unmarshal(marshalMap(),&umdata)
	if err!=nil{
		log.Fatal("Error in unmarshaling: ",err)
	}
	fmt.Println(umdata)
}
