package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

type Message struct {
	Userid int
	MsgId int `json:"id"`
	Title,Body string
}

const url  ="https://jsonplaceholder.typicode.com/posts/1";

var m =Message{Userid:1}
func main()  {
	resp,err:=http.Get(url)
	if err!=nil && resp.StatusCode!=200{
		resp.Body.Close()
		log.Fatal("Error in request: ",err)
	}
	if  resp.StatusCode!=200{
		resp.Body.Close()
		log.Fatal("Error With request: ",resp.Status)
	}
	defer resp.Body.Close()
	err=json.NewDecoder(resp.Body).Decode(&m)
	if err!=nil{
		log.Fatal("Error in decoding: ",err)
	}

	fmt.Println(m.Userid,"\n",m.MsgId,"\n",m.Title+"\n",m.Body,"\n")

}

