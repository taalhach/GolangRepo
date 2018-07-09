package main

import (
	"encoding/json"
	"os"
	"log"
)

type Message struct {
	Userid string
	MsgId int `json:"id"`
	Title,Body string
}
var m = Message{
	Userid:"taalhach",
	MsgId:123,
	Title:"Important task needs to be done immediately",
	Body:"You are needed here in Election commission and a task will be assigned to  you so that you can do it quickly",
}

func main()  {

	err:=json.NewEncoder(os.Stdout).Encode(&m)
	if err!=nil{
		log.Fatal("Error in encoding: ",err)
	}
}

